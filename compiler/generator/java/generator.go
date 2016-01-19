package java

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Workiva/frugal/compiler/generator"
	"github.com/Workiva/frugal/compiler/globals"
	"github.com/Workiva/frugal/compiler/parser"
)

const (
	lang                     = "java"
	defaultOutputDir         = "gen-java"
	tab                      = "\t"
	tabtab                   = tab + tab
	tabtabtab                = tab + tab + tab
	tabtabtabtab             = tab + tab + tab + tab
	tabtabtabtabtab          = tab + tab + tab + tab + tab
	tabtabtabtabtabtab       = tab + tab + tab + tab + tab + tab
	tabtabtabtabtabtabtab    = tab + tab + tab + tab + tab + tab + tab
	tabtabtabtabtabtabtabtab = tab + tab + tab + tab + tab + tab + tab + tab
)

type Generator struct {
	*generator.BaseGenerator
	time time.Time
}

func NewGenerator(options map[string]string) generator.LanguageGenerator {
	return &Generator{
		&generator.BaseGenerator{Options: options},
		globals.Now,
	}
}

func (g *Generator) GenerateThrift() bool {
	return true
}

func (g *Generator) GetOutputDir(dir string) string {
	if pkg, ok := g.Frugal.Thrift.Namespace(lang); ok {
		path := generator.GetPackageComponents(pkg)
		dir = filepath.Join(append([]string{dir}, path...)...)
	}
	return dir
}

func (g *Generator) DefaultOutputDir() string {
	return defaultOutputDir
}

func (g *Generator) GenerateDependencies(dir string) error {
	return nil
}

func (g *Generator) GenerateFile(name, outputDir string, fileType generator.FileType) (*os.File, error) {
	switch fileType {
	case generator.PublishFile:
		return g.CreateFile(strings.Title(name)+"Publisher", outputDir, lang, false)
	case generator.SubscribeFile:
		return g.CreateFile(strings.Title(name)+"Subscriber", outputDir, lang, false)
	case generator.CombinedServiceFile:
		return g.CreateFile("F"+name, outputDir, lang, false)
	default:
		return nil, fmt.Errorf("frugal: Bad file type for Java generator: %s", fileType)
	}
}

func (g *Generator) GenerateDocStringComment(file *os.File) error {
	comment := fmt.Sprintf(
		"/**\n"+
			" * Autogenerated by Frugal Compiler (%s)\n"+
			" * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING\n"+
			" *  @generated\n"+
			" */",
		globals.Version)

	_, err := file.WriteString(comment)
	return err
}

func (g *Generator) GenerateServicePackage(file *os.File, s *parser.Service) error {
	return g.generatePackage(file)
}

func (g *Generator) GenerateScopePackage(file *os.File, s *parser.Scope) error {
	return g.generatePackage(file)
}

func (g *Generator) generatePackage(file *os.File) error {
	pkg, ok := g.Frugal.Thrift.Namespace(lang)
	if !ok {
		return nil
	}
	_, err := file.WriteString(fmt.Sprintf("package %s;", pkg))
	return err
}

func (g *Generator) GenerateServiceImports(file *os.File, s *parser.Service) error {
	imports := "import com.workiva.frugal.*;\n"
	imports += "import com.workiva.frugal.processor.FProcessor;\n"
	imports += "import com.workiva.frugal.processor.FProcessorFunction;\n"
	imports += "import com.workiva.frugal.registry.FAsyncCallback;\n"
	imports += "import com.workiva.frugal.registry.FClientRegistry;\n"
	imports += "import com.workiva.frugal.transport.FTransport;\n"
	imports += "import org.apache.thrift.TApplicationException;\n"
	imports += "import org.apache.thrift.TException;\n"
	imports += "import org.apache.thrift.protocol.TMessage;\n"
	imports += "import org.apache.thrift.protocol.TMessageType;\n"
	imports += "import org.apache.thrift.protocol.TProtocolUtil;\n"
	imports += "import org.apache.thrift.protocol.TType;\n"
	imports += "import org.apache.thrift.transport.TTransport;\n\n"

	imports += "import javax.annotation.Generated;\n"
	imports += "import java.util.HashMap;\n"
	imports += "import java.util.Map;\n"
	imports += "import java.util.concurrent.BlockingQueue;\n"
	imports += "import java.util.concurrent.ArrayBlockingQueue;\n"
	imports += "import java.util.concurrent.TimeUnit;\n"

	_, err := file.WriteString(imports)
	return err
}

func (g *Generator) GenerateScopeImports(file *os.File, s *parser.Scope) error {
	imports := "import com.workiva.frugal.FContext;\n"
	imports += "import com.workiva.frugal.FScopeProvider;\n"
	imports += "import com.workiva.frugal.FSubscription;\n"
	imports += "import com.workiva.frugal.FProtocol;\n"
	imports += "import com.workiva.frugal.transport.FScopeTransport;\n"
	imports += "import org.apache.thrift.TException;\n"
	imports += "import org.apache.thrift.TApplicationException;\n"
	imports += "import org.apache.thrift.transport.TTransportException;\n"
	imports += "import org.apache.thrift.protocol.*;\n\n"

	imports += "import javax.annotation.Generated;\n"
	imports += "import java.util.logging.Logger;\n"

	_, err := file.WriteString(imports)
	return err
}

func (g *Generator) GenerateConstants(file *os.File, name string) error {
	return nil
}

func (g *Generator) GeneratePublisher(file *os.File, scope *parser.Scope) error {
	publisher := ""
	if scope.Comment != nil {
		publisher += g.GenerateBlockComment(scope.Comment, "")
	}
	publisher += fmt.Sprintf("@Generated(value = \"Autogenerated by Frugal Compiler (%s)\", "+
		"date = \"%s\")\n", globals.Version, g.time.Format("2006-1-2"))
	publisher += fmt.Sprintf("public class %sPublisher {\n\n", strings.Title(scope.Name))

	publisher += fmt.Sprintf(tab+"private static final String DELIMITER = \"%s\";\n\n", globals.TopicDelimiter)

	publisher += tab + "private final FScopeProvider provider;\n"
	publisher += tab + "private FScopeTransport transport;\n"
	publisher += tab + "private FProtocol protocol;\n\n"

	publisher += fmt.Sprintf(tab+"public %sPublisher(FScopeProvider provider) {\n", strings.Title(scope.Name))
	publisher += tabtab + "this.provider = provider;\n"
	publisher += tab + "}\n\n"

	publisher += tab + "public void open() throws TException {\n"
	publisher += tabtab + "FScopeProvider.Client client = provider.build();\n"
	publisher += tabtab + "transport = client.getTransport();\n"
	publisher += tabtab + "protocol = client.getProtocol();\n"
	publisher += tabtab + "transport.open();\n"
	publisher += tab + "}\n\n"

	publisher += tab + "public void close() throws TException {\n"
	publisher += tabtab + "transport.close();\n"
	publisher += tab + "}\n\n"

	args := ""
	if len(scope.Prefix.Variables) > 0 {
		for _, variable := range scope.Prefix.Variables {
			args = fmt.Sprintf("%sString %s, ", args, variable)
		}
	}
	prefix := ""
	for _, op := range scope.Operations {
		publisher += prefix
		prefix = "\n\n"
		if op.Comment != nil {
			publisher += g.GenerateBlockComment(op.Comment, tab)
		}
		publisher += fmt.Sprintf(tab+"public void publish%s(FContext ctx, %s%s req) throws TException {\n", op.Name, args, g.qualifiedParamName(op))
		publisher += fmt.Sprintf(tabtab+"String op = \"%s\";\n", op.Name)
		publisher += fmt.Sprintf(tabtab+"String prefix = %s;\n", generatePrefixStringTemplate(scope))
		publisher += tabtab + "String topic = String.format(\"%s" + strings.Title(scope.Name) + "%s%s\", prefix, DELIMITER, op);\n"
		publisher += tabtab + "transport.lockTopic(topic);\n"
		publisher += tabtab + "try {\n"
		publisher += tabtabtab + "protocol.writeRequestHeader(ctx);\n"
		publisher += tabtabtab + "protocol.writeMessageBegin(new TMessage(op, TMessageType.CALL, 0));\n"
		publisher += tabtabtab + "req.write(protocol);\n"
		publisher += tabtabtab + "protocol.writeMessageEnd();\n"
		publisher += tabtabtab + "transport.flush();\n"
		publisher += tabtab + "} catch (TException e) {\n"
		publisher += tabtabtab + "close();\n"
		publisher += tabtabtab + "throw e;\n"
		publisher += tabtab + "} finally {\n"
		publisher += tabtabtab + "transport.unlockTopic();\n"
		publisher += tabtab + "}\n"
		publisher += tab + "}\n"
	}

	publisher += "}"

	_, err := file.WriteString(publisher)
	return err
}

func generatePrefixStringTemplate(scope *parser.Scope) string {
	if len(scope.Prefix.Variables) == 0 {
		if scope.Prefix.String == "" {
			return `""`
		}
		return fmt.Sprintf(`"%s%s"`, scope.Prefix.String, globals.TopicDelimiter)
	}
	template := "String.format(\""
	template += scope.Prefix.Template()
	template += globals.TopicDelimiter + "\", "
	prefix := ""
	for _, variable := range scope.Prefix.Variables {
		template += prefix + variable
		prefix = ", "
	}
	template += ")"
	return template
}

func (g *Generator) GenerateSubscriber(file *os.File, scope *parser.Scope) error {
	subscriber := ""
	if scope.Comment != nil {
		subscriber += g.GenerateBlockComment(scope.Comment, "")
	}
	scopeName := strings.Title(scope.Name)
	subscriber += fmt.Sprintf("@Generated(value = \"Autogenerated by Frugal Compiler (%s)\", "+
		"date = \"%s\")\n", globals.Version, g.time.Format("2006-1-2"))
	subscriber += fmt.Sprintf("public class %sSubscriber {\n\n", scopeName)

	subscriber += fmt.Sprintf(tab+"private static final String DELIMITER = \"%s\";\n", globals.TopicDelimiter)
	subscriber += fmt.Sprintf(
		tab+"private static Logger LOGGER = Logger.getLogger(%sSubscriber.class.getName());\n\n", scopeName)

	subscriber += tab + "private final FScopeProvider provider;\n\n"

	subscriber += fmt.Sprintf(tab+"public %sSubscriber(FScopeProvider provider) {\n",
		strings.Title(scope.Name))
	subscriber += tabtab + "this.provider = provider;\n"
	subscriber += tab + "}\n\n"

	args := ""
	if len(scope.Prefix.Variables) > 0 {
		for _, variable := range scope.Prefix.Variables {
			args = fmt.Sprintf("%sString %s, ", args, variable)
		}
	}
	prefix := ""
	for _, op := range scope.Operations {
		subscriber += fmt.Sprintf(tab+"public interface %sHandler {\n", op.Name)
		subscriber += fmt.Sprintf(tabtab+"void on%s(FContext ctx, %s req);\n", op.Name, g.qualifiedParamName(op))
		subscriber += tab + "}\n\n"

		subscriber += prefix
		prefix = "\n\n"
		if op.Comment != nil {
			subscriber += g.GenerateBlockComment(op.Comment, tab)
		}
		subscriber += fmt.Sprintf(tab+"public FSubscription subscribe%s(%sfinal %sHandler handler) throws TException {\n",
			op.Name, args, op.Name)
		subscriber += fmt.Sprintf(tabtab+"final String op = \"%s\";\n", op.Name)
		subscriber += fmt.Sprintf(tabtab+"String prefix = %s;\n", generatePrefixStringTemplate(scope))
		subscriber += tabtab + "String topic = String.format(\"%s" + strings.Title(scope.Name) + "%s%s\", prefix, DELIMITER, op);\n"
		subscriber += tabtab + "final FScopeProvider.Client client = provider.build();\n"
		subscriber += tabtab + "FScopeTransport transport = client.getTransport();\n"
		subscriber += tabtab + "transport.subscribe(topic);\n\n"

		subscriber += tabtab + "final FSubscription sub = new FSubscription(topic, transport);\n"
		subscriber += tabtab + "new Thread(new Runnable() {\n"
		subscriber += tabtabtab + "public void run() {\n"
		subscriber += tabtabtabtab + "while (true) {\n"
		subscriber += tabtabtabtabtab + "try {\n"
		subscriber += tabtabtabtabtabtab + "FContext ctx = client.getProtocol().readRequestHeader();\n"
		subscriber += tabtabtabtabtabtab + fmt.Sprintf("%s received = recv%s(op, client.getProtocol());\n",
			g.qualifiedParamName(op), op.Name)
		subscriber += tabtabtabtabtabtab + fmt.Sprintf("handler.on%s(ctx, received);\n", op.Name)
		subscriber += tabtabtabtabtab + "} catch (TException e) {\n"
		subscriber += tabtabtabtabtabtab + "if (e instanceof TTransportException) {\n"
		subscriber += tabtabtabtabtabtabtab + "TTransportException transportException = (TTransportException) e;\n"
		subscriber += tabtabtabtabtabtabtab + "if (transportException.getType() == TTransportException.END_OF_FILE) {\n"
		subscriber += tabtabtabtabtabtabtabtab + "return;\n"
		subscriber += tabtabtabtabtabtabtab + "}\n"
		subscriber += tabtabtabtabtabtab + "}\n"
		subscriber += tabtabtabtabtabtab + fmt.Sprintf("LOGGER.severe(\"Subscriber recv%s error \" + e.getMessage());\n", op.Name)
		subscriber += tabtabtabtabtabtab + "sub.signal(e);\n"
		subscriber += tabtabtabtabtabtab + "sub.unsubscribe();\n"
		subscriber += tabtabtabtabtabtab + "return;\n"
		subscriber += tabtabtabtabtab + "}\n"
		subscriber += tabtabtabtab + "}\n"
		subscriber += tabtabtab + "}\n"
		subscriber += tabtab + "}).start();\n\n"

		subscriber += tabtab + "return sub;\n"
		subscriber += tab + "}\n\n"

		subscriber += tab + fmt.Sprintf("private %s recv%s(String op, FProtocol iprot) throws TException {\n", g.qualifiedParamName(op), op.Name)
		subscriber += tabtab + "TMessage msg = iprot.readMessageBegin();\n"
		subscriber += tabtab + "if (!msg.name.equals(op)) {\n"
		subscriber += tabtabtab + "TProtocolUtil.skip(iprot, TType.STRUCT);\n"
		subscriber += tabtabtab + "iprot.readMessageEnd();\n"
		subscriber += tabtabtab + "throw new TApplicationException(TApplicationException.UNKNOWN_METHOD);\n"
		subscriber += tabtab + "}\n"
		subscriber += tabtab + fmt.Sprintf("%s req = new %s();\n", g.qualifiedParamName(op), g.qualifiedParamName(op))
		subscriber += tabtab + "req.read(iprot);\n"
		subscriber += tabtab + "iprot.readMessageEnd();\n"
		subscriber += tabtab + "return req;\n"
		subscriber += tab + "}\n\n"
	}
	subscriber += "\n}"

	_, err := file.WriteString(subscriber)
	return err
}

func (g *Generator) GenerateService(file *os.File, s *parser.Service) error {
	contents := ""
	contents += fmt.Sprintf("@Generated(value = \"Autogenerated by Frugal Compiler (%s)\", "+
		"date = \"%s\")\n", globals.Version, g.time.Format("2006-1-2"))
	contents += fmt.Sprintf("public class F%s {\n\n", s.Name)
	contents += g.generateServiceInterface(s)
	contents += g.generateClient(s)
	contents += g.generateServer(s)

	_, err := file.WriteString(contents)
	return err
}

func (g *Generator) generateServiceInterface(service *parser.Service) string {
	contents := ""
	if service.Comment != nil {
		contents += g.GenerateBlockComment(service.Comment, tab)
	}
	contents += tab + "public interface Iface {\n\n"
	for _, method := range service.Methods {
		if method.Comment != nil {
			contents += g.GenerateBlockComment(method.Comment, tabtab)
		}
		contents += fmt.Sprintf(tabtab+"public %s %s(FContext ctx%s) %s;\n\n",
			g.generateReturnValue(method), method.Name, g.generateArgs(method.Arguments), g.generateExceptions(method.Exceptions))
	}
	contents += "}\n\n"
	return contents
}

func (g *Generator) generateReturnValue(method *parser.Method) string {
	if method.ReturnType == nil {
		return "void"
	}
	return g.getJavaTypeFromThriftType(method.ReturnType)
}

func (g *Generator) generateArgs(args []*parser.Field) string {
	argStr := ""
	for _, arg := range args {
		argStr += ", " + g.getJavaTypeFromThriftType(arg.Type) + " " + arg.Name
	}
	return argStr
}

func (g *Generator) generateClient(service *parser.Service) string {
	contents := tab + "public static class Client implements Iface {\n\n"
	contents += tabtab + "private static final Object WRITE_LOCK = new Object();\n\n"

	contents += tabtab + "private FTransport transport;\n"
	contents += tabtab + "private FProtocolFactory protocolFactory;\n"
	contents += tabtab + "private FProtocol inputProtocol;\n"
	contents += tabtab + "private FProtocol outputProtocol;\n\n"

	contents += tabtab + "public Client(FServiceProvider provider) {\n"
	contents += tabtabtab + "this.transport = provider.getTransport();\n"
	contents += tabtabtab + "this.transport.setRegistry(new FClientRegistry());\n"
	contents += tabtabtab + "this.protocolFactory = provider.getProtocolFactory();\n"
	contents += tabtabtab + "this.inputProtocol = this.protocolFactory.getProtocol(this.transport);\n"
	contents += tabtabtab + "this.outputProtocol = this.protocolFactory.getProtocol(this.transport);\n"
	contents += tabtab + "}\n\n"

	for _, method := range service.Methods {
		contents += g.generateClientMethod(service, method)
	}
	contents += tab + "}\n\n"

	return contents
}

func (g *Generator) generateClientMethod(service *parser.Service, method *parser.Method) string {
	servTitle := strings.Title(service.Name)

	contents := ""
	if method.Comment != nil {
		contents += g.GenerateBlockComment(method.Comment, tabtab)
	}
	contents += tabtab + fmt.Sprintf("public %s %s(FContext ctx%s) %s {\n",
		g.generateReturnValue(method), method.Name, g.generateArgs(method.Arguments), g.generateExceptions(method.Exceptions))
	contents += tabtabtab + "FProtocol oprot = this.outputProtocol;\n"
	contents += tabtabtab + "BlockingQueue<Object> result = new ArrayBlockingQueue<>(1);\n"
	contents += tabtabtab + fmt.Sprintf("this.transport.register(ctx, recv%sHandler(ctx, result));\n", strings.Title(method.Name))
	contents += tabtabtab + "try {\n"
	contents += tabtabtabtab + "synchronized (WRITE_LOCK) {\n"
	contents += tabtabtabtabtab + "oprot.writeRequestHeader(ctx);\n"
	contents += tabtabtabtabtab + fmt.Sprintf("oprot.writeMessageBegin(new TMessage(\"%s\", TMessageType.CALL, 0));\n", method.Name)
	contents += tabtabtabtabtab + fmt.Sprintf("%s.%s_args args = new %s.%s_args();\n", servTitle, method.Name, servTitle, method.Name)
	for _, arg := range method.Arguments {
		contents += tabtabtabtabtab + fmt.Sprintf("args.set%s(%s);\n", strings.Title(arg.Name), arg.Name)
	}
	contents += tabtabtabtabtab + "args.write(oprot);\n"
	contents += tabtabtabtabtab + "oprot.writeMessageEnd();\n"
	contents += tabtabtabtabtab + "oprot.getTransport().flush();\n"
	contents += tabtabtabtab + "}\n\n"

	contents += tabtabtabtab + "Object res = null;\n"
	contents += tabtabtabtab + "try {\n"
	contents += tabtabtabtabtab + "res = result.poll(ctx.getTimeout(), TimeUnit.MILLISECONDS);\n"
	contents += tabtabtabtab + "} catch (InterruptedException e) {\n"
	contents += tabtabtabtabtab + fmt.Sprintf(
		"throw new TApplicationException(TApplicationException.INTERNAL_ERROR, \"%s interrupted: \" + e.getMessage());\n",
		method.Name)
	contents += tabtabtabtab + "}\n"
	contents += tabtabtabtab + "if (res == null) {\n"
	contents += tabtabtabtabtab + fmt.Sprintf("throw new FTimeoutException(\"%s timed out\");\n", method.Name)
	contents += tabtabtabtab + "}\n"
	contents += tabtabtabtab + "if (res instanceof TException) {\n"
	contents += tabtabtabtabtab + "throw (TException) res;\n"
	contents += tabtabtabtab + "}\n"
	contents += tabtabtabtab + fmt.Sprintf("%s.%s_result r = (%s.%s_result) res;\n", servTitle, method.Name, servTitle, method.Name)
	if method.ReturnType != nil {
		contents += tabtabtabtab + "if (r.isSetSuccess()) {\n"
		contents += tabtabtabtabtab + "return r.success;\n"
		contents += tabtabtabtab + "}\n"
	}
	for _, exception := range method.Exceptions {
		contents += tabtabtabtab + fmt.Sprintf("if (r.%s != null) {\n", exception.Name)
		contents += tabtabtabtabtab + fmt.Sprintf("throw r.%s;\n", exception.Name)
		contents += tabtabtabtab + "}\n"
	}
	if method.ReturnType != nil {
		contents += tabtabtabtab + fmt.Sprintf(
			"throw new TApplicationException(TApplicationException.MISSING_RESULT, \"%s failed: unknown result\");\n",
			method.Name)
	}
	contents += tabtabtab + "} finally {\n"
	contents += tabtabtabtab + "this.transport.unregister(ctx);\n"
	contents += tabtabtab + "}\n"
	contents += tabtab + "}\n\n"

	contents += tabtab + fmt.Sprintf(
		"private FAsyncCallback recv%sHandler(final FContext ctx, final BlockingQueue<Object> result) {\n",
		strings.Title(method.Name))
	contents += tabtabtab + "return new FAsyncCallback() {\n"
	contents += tabtabtabtab + "public void onMessage(TTransport tr) throws TException {\n"
	contents += tabtabtabtabtab + "FProtocol iprot = Client.this.protocolFactory.getProtocol(tr);\n"
	contents += tabtabtabtabtab + "try {\n"
	contents += tabtabtabtabtabtab + "iprot.readResponseHeader(ctx);\n"
	contents += tabtabtabtabtabtab + "TMessage message = iprot.readMessageBegin();\n"
	contents += tabtabtabtabtabtab + fmt.Sprintf("if (!message.name.equals(\"%s\")) {\n", method.Name)
	contents += tabtabtabtabtabtabtab + fmt.Sprintf(
		"throw new TApplicationException(TApplicationException.WRONG_METHOD_NAME, \"%s failed: wrong method name\");\n",
		method.Name)
	contents += tabtabtabtabtabtab + "}\n"
	contents += tabtabtabtabtabtab + "if (message.type == TMessageType.EXCEPTION) {\n"
	contents += tabtabtabtabtabtabtab + "TApplicationException e = TApplicationException.read(iprot);\n"
	contents += tabtabtabtabtabtabtab + "iprot.readMessageEnd();\n"
	contents += tabtabtabtabtabtabtab + "throw e;\n"
	contents += tabtabtabtabtabtab + "}\n"
	contents += tabtabtabtabtabtab + "if (message.type != TMessageType.REPLY) {\n"
	contents += tabtabtabtabtabtabtab + fmt.Sprintf(
		"throw new TApplicationException(TApplicationException.INVALID_MESSAGE_TYPE, \"%s failed: invalid message type\");\n",
		method.Name)
	contents += tabtabtabtabtabtab + "}\n"
	contents += tabtabtabtabtabtab + fmt.Sprintf("%s.%s_result res = new %s.%s_result();\n", servTitle, method.Name, servTitle, method.Name)
	contents += tabtabtabtabtabtab + "res.read(iprot);\n"
	contents += tabtabtabtabtabtab + "iprot.readMessageEnd();\n"
	contents += tabtabtabtabtabtab + "try {\n"
	contents += tabtabtabtabtabtabtab + "result.put(res);\n"
	contents += tabtabtabtabtabtab + "} catch (InterruptedException e) {\n"
	contents += tabtabtabtabtabtabtab + fmt.Sprintf(
		"throw new TApplicationException(TApplicationException.INTERNAL_ERROR, \"%s interrupted: \" + e.getMessage());\n",
		method.Name)
	contents += tabtabtabtabtabtab + "}\n"
	contents += tabtabtabtabtab + "} catch (TException e) {\n"
	contents += tabtabtabtabtabtab + "try {\n"
	contents += tabtabtabtabtabtabtab + "result.put(e);\n"
	contents += tabtabtabtabtabtab + "} finally {\n"
	contents += tabtabtabtabtabtabtab + "throw e;\n"
	contents += tabtabtabtabtabtab + "}\n"
	contents += tabtabtabtabtab + "}\n"
	contents += tabtabtabtab + "}\n"
	contents += tabtabtab + "};\n"
	contents += tabtab + "}\n\n"

	return contents
}

func (g *Generator) generateExceptions(exceptions []*parser.Field) string {
	contents := "throws TException"
	for _, exception := range exceptions {
		contents += ", " + exception.Type.String()
	}
	return contents
}

func (g *Generator) generateServer(service *parser.Service) string {
	servTitle := strings.Title(service.Name)

	contents := ""
	contents += tab + "public static class Processor implements FProcessor {\n\n"

	contents += tabtab + "private static final Object WRITE_LOCK = new Object();\n\n"

	contents += tabtab + "private Map<String, FProcessorFunction> processorMap = new HashMap<>();\n\n"

	contents += tabtab + "public Processor(Iface handler) {\n"
	for _, method := range service.Methods {
		contents += tabtabtab + fmt.Sprintf("this.processorMap.put(\"%s\", new %s(handler));\n", method.Name, strings.Title(method.Name))
	}
	contents += tabtab + "}\n\n"

	contents += tabtab + "public void process(FProtocol iprot, FProtocol oprot) throws TException {\n"
	contents += tabtabtab + "FContext ctx = iprot.readRequestHeader();\n"
	contents += tabtabtab + "TMessage message = iprot.readMessageBegin();\n"
	contents += tabtabtab + "FProcessorFunction processor = this.processorMap.get(message.name);\n"
	contents += tabtabtab + "if (processor != null) {\n"
	contents += tabtabtabtab + "processor.process(ctx, iprot, oprot);\n"
	contents += tabtabtabtab + "return;\n"
	contents += tabtabtab + "}\n"
	contents += tabtabtab + "TProtocolUtil.skip(iprot, TType.STRUCT);\n"
	contents += tabtabtab + "iprot.readMessageEnd();\n"
	contents += tabtabtab + "TApplicationException e = new TApplicationException(TApplicationException.UNKNOWN_METHOD, \"Unknown function \" + message.name);\n"
	contents += tabtabtab + "synchronized (WRITE_LOCK) {\n"
	contents += tabtabtabtab + "oprot.writeResponseHeader(ctx);\n"
	contents += tabtabtabtab + "oprot.writeMessageBegin(new TMessage(message.name, TMessageType.EXCEPTION, 0));\n"
	contents += tabtabtabtab + "e.write(oprot);\n"
	contents += tabtabtabtab + "oprot.writeMessageEnd();\n"
	contents += tabtabtabtab + "oprot.getTransport().flush();\n"
	contents += tabtabtab + "}\n"
	contents += tabtabtab + "throw e;\n"
	contents += tabtab + "}\n\n"

	for _, method := range service.Methods {
		contents += tabtab + fmt.Sprintf("private static class %s implements FProcessorFunction {\n\n", strings.Title(method.Name))

		contents += tabtabtab + "private Iface handler;\n\n"

		contents += tabtabtab + fmt.Sprintf("public %s(Iface handler) {\n", strings.Title(method.Name))
		contents += tabtabtabtab + "this.handler = handler;\n"
		contents += tabtabtab + "}\n\n"

		contents += tabtabtab + "public void process(FContext ctx, FProtocol iprot, FProtocol oprot) throws TException {\n"
		contents += tabtabtabtab + fmt.Sprintf("%s.%s_args args = new %s.%s_args();\n", servTitle, method.Name, servTitle, method.Name)
		contents += tabtabtabtab + "try {\n"
		contents += tabtabtabtabtab + "args.read(iprot);\n"
		contents += tabtabtabtab + "} catch (TException e) {\n"
		contents += tabtabtabtabtab + "iprot.readMessageEnd();\n"
		contents += tabtabtabtabtab + "TApplicationException x = new TApplicationException(TApplicationException.PROTOCOL_ERROR, e.getMessage());\n"
		contents += tabtabtabtabtab + "synchronized (WRITE_LOCK) {\n"
		contents += tabtabtabtabtabtab + "oprot.writeResponseHeader(ctx);\n"
		contents += tabtabtabtabtabtab + fmt.Sprintf("oprot.writeMessageBegin(new TMessage(\"%s\", TMessageType.EXCEPTION, 0));\n", method.Name)
		contents += tabtabtabtabtabtab + "x.write(oprot);\n"
		contents += tabtabtabtabtabtab + "oprot.writeMessageEnd();\n"
		contents += tabtabtabtabtabtab + "oprot.getTransport().flush();\n"
		contents += tabtabtabtabtab + "}\n"
		contents += tabtabtabtabtab + "throw x;\n"
		contents += tabtabtabtab + "}\n\n"

		contents += tabtabtabtab + "iprot.readMessageEnd();\n"
		contents += tabtabtabtab + fmt.Sprintf("%s.%s_result result = new %s.%s_result();\n", servTitle, method.Name, servTitle, method.Name)
		contents += tabtabtabtab + "try {\n"
		if method.ReturnType == nil {
			contents += tabtabtabtabtab + fmt.Sprintf("this.handler.%s(%s);\n", method.Name, g.generateCallArgs(method.Arguments))
		} else {
			contents += tabtabtabtabtab + fmt.Sprintf("result.success = this.handler.%s(%s);\n", method.Name, g.generateCallArgs(method.Arguments))
			contents += tabtabtabtabtab + "result.setSuccessIsSet(true);\n"
		}
		for _, exception := range method.Exceptions {
			contents += tabtabtabtab + fmt.Sprintf("} catch (%s %s) {\n", exception.Type, exception.Name)
			contents += tabtabtabtabtab + fmt.Sprintf("result.%s = %s;\n", exception.Name, exception.Name)
		}
		contents += tabtabtabtab + "} catch (TException e) {\n"
		contents += tabtabtabtabtab + fmt.Sprintf(
			"TApplicationException x = new TApplicationException(TApplicationException.INTERNAL_ERROR, \"Internal error processing %s: \" + e.getMessage());\n",
			method.Name)
		contents += tabtabtabtabtab + "synchronized (WRITE_LOCK) {\n"
		contents += tabtabtabtabtabtab + "oprot.writeResponseHeader(ctx);\n"
		contents += tabtabtabtabtabtab + fmt.Sprintf("oprot.writeMessageBegin(new TMessage(\"%s\", TMessageType.EXCEPTION, 0));\n", method.Name)
		contents += tabtabtabtabtabtab + "x.write(oprot);\n"
		contents += tabtabtabtabtabtab + "oprot.writeMessageEnd();\n"
		contents += tabtabtabtabtabtab + "oprot.getTransport().flush();\n"
		contents += tabtabtabtabtab + "}\n"
		contents += tabtabtabtabtab + "throw e;\n"
		contents += tabtabtabtab + "}\n"
		contents += tabtabtabtab + "synchronized (WRITE_LOCK) {\n"
		contents += tabtabtabtabtab + "oprot.writeResponseHeader(ctx);\n"
		contents += tabtabtabtabtab + fmt.Sprintf("oprot.writeMessageBegin(new TMessage(\"%s\", TMessageType.REPLY, 0));\n", method.Name)
		contents += tabtabtabtabtab + "result.write(oprot);\n"
		contents += tabtabtabtabtab + "oprot.writeMessageEnd();\n"
		contents += tabtabtabtabtab + "oprot.getTransport().flush();\n"
		contents += tabtabtabtab + "}\n"
		contents += tabtabtab + "}\n"
		contents += tabtab + "}\n\n"
	}
	contents += tab + "}\n\n"
	contents += "}"

	return contents
}

func (g *Generator) generateCallArgs(args []*parser.Field) string {
	contents := "ctx"
	for _, arg := range args {
		contents += ", args." + arg.Name
	}
	return contents
}

func (g *Generator) getJavaTypeFromThriftType(t *parser.Type) string {
	if t == nil {
		return "void"
	}
	typeName := g.Frugal.Thrift.UnderlyingType(t.Name)
	switch typeName {
	case "bool":
		return "boolean"
	case "byte":
		return "byte"
	case "i16":
		return "short"
	case "i32":
		return "int"
	case "i64":
		return "long"
	case "double":
		return "double"
	case "string":
		return "String"
	case "binary":
		return "java.nio.ByteBuffer"
	case "list":
		return fmt.Sprintf("List<%s>", g.getJavaTypeFromThriftType(t.ValueType))
	case "set":
		return fmt.Sprintf("Set<%s>", g.getJavaTypeFromThriftType(t.ValueType))
	case "map":
		return fmt.Sprintf("Map<%s, %s>", g.getJavaTypeFromThriftType(t.KeyType),
			g.getJavaTypeFromThriftType(t.ValueType))
	default:
		// This is a custom type, return a pointer to it
		return g.qualifiedTypeName(t)
	}
}

func (g *Generator) qualifiedTypeName(t *parser.Type) string {
	param := t.ParamName()
	include := t.IncludeName()
	if include != "" {
		namespace, ok := g.Frugal.NamespaceForInclude(include, lang)
		if !ok {
			namespace = include
		}
		param = fmt.Sprintf("%s.%s", namespace, param)
	}
	return param
}

func (g *Generator) qualifiedParamName(op *parser.Operation) string {
	param := op.ParamName()
	include := op.IncludeName()
	if include != "" {
		namespace, ok := g.Frugal.NamespaceForInclude(include, lang)
		if ok {
			param = fmt.Sprintf("%s.%s", namespace, param)
		}
	}
	return param
}
