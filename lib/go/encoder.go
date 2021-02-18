package frugal

import "github.com/apache/thrift/lib/go/thrift"

// WriteString writes string `value` of field name and id `name` and `field` respectively into `p`.
func WriteString(fctx FContext, p thrift.TProtocol, value, name string, field int16) error {
	ctx := toCTX(fctx)
	if err := p.WriteFieldBegin(ctx, name, thrift.STRING, field); err != nil {
		return thrift.PrependError("write field begin error: ", err)
	}
	if err := p.WriteString(ctx, value); err != nil {
		return thrift.PrependError("field write error: ", err)
	}
	if err := p.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError("write field end error: ", err)
	}
	return nil
}

// WriteBool writes bool `value` of field name and id `name` and `field` respectively into `p`.
func WriteBool(fctx FContext, p thrift.TProtocol, value bool, name string, field int16) error {
	ctx := toCTX(fctx)
	if err := p.WriteFieldBegin(ctx, name, thrift.BOOL, field); err != nil {
		return thrift.PrependError("write field begin error: ", err)
	}
	if err := p.WriteBool(ctx, value); err != nil {
		return thrift.PrependError("field write error: ", err)
	}
	if err := p.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError("write field end error: ", err)
	}
	return nil
}

// WriteByte writes byte `value` of field name and id `name` and `field` respectively into `p`.
func WriteByte(fctx FContext, p thrift.TProtocol, value int8, name string, field int16) error {
	ctx := toCTX(fctx)
	if err := p.WriteFieldBegin(ctx, name, thrift.BYTE, field); err != nil {
		return thrift.PrependError("write field begin error: ", err)
	}
	if err := p.WriteByte(ctx, value); err != nil {
		return thrift.PrependError("field write error: ", err)
	}
	if err := p.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError("write field end error: ", err)
	}
	return nil
}

// WriteDouble writes float64 `value` of field name and id `name` and `field` respectively into `p`.
func WriteDouble(fctx FContext, p thrift.TProtocol, value float64, name string, field int16) error {
	ctx := toCTX(fctx)
	if err := p.WriteFieldBegin(ctx, name, thrift.DOUBLE, field); err != nil {
		return thrift.PrependError("write field begin error: ", err)
	}
	if err := p.WriteDouble(ctx, value); err != nil {
		return thrift.PrependError("field write error: ", err)
	}
	if err := p.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError("write field end error: ", err)
	}
	return nil
}

// WriteI16 writes int16 `value` of field name and id `name` and `field` respectively into `p`.
func WriteI16(fctx FContext, p thrift.TProtocol, value int16, name string, field int16) error {
	ctx := toCTX(fctx)
	if err := p.WriteFieldBegin(ctx, name, thrift.I16, field); err != nil {
		return thrift.PrependError("write field begin error: ", err)
	}
	if err := p.WriteI16(ctx, value); err != nil {
		return thrift.PrependError("field write error: ", err)
	}
	if err := p.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError("write field end error: ", err)
	}
	return nil
}

// WriteI32 writes int32 `value` of field name and id `name` and `field` respectively into `p`.
func WriteI32(fctx FContext, p thrift.TProtocol, value int32, name string, field int16) error {
	ctx := toCTX(fctx)
	if err := p.WriteFieldBegin(ctx, name, thrift.I32, field); err != nil {
		return thrift.PrependError("write field begin error: ", err)
	}
	if err := p.WriteI32(ctx, value); err != nil {
		return thrift.PrependError("field write error: ", err)
	}
	if err := p.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError("write field end error: ", err)
	}
	return nil
}

// WriteI64 writes int64 `value` of field name and id `name` and `field` respectively into `p`.
func WriteI64(fctx FContext, p thrift.TProtocol, value int64, name string, field int16) error {
	ctx := toCTX(fctx)
	if err := p.WriteFieldBegin(ctx, name, thrift.I64, field); err != nil {
		return thrift.PrependError("write field begin error: ", err)
	}
	if err := p.WriteI64(ctx, value); err != nil {
		return thrift.PrependError("field write error: ", err)
	}
	if err := p.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError("write field end error: ", err)
	}
	return nil
}

// WriteBinary writes []byte `value` of field name and id `name` and `field` respectively into `p`.
func WriteBinary(fctx FContext, p thrift.TProtocol, value []byte, name string, field int16) error {
	ctx := toCTX(fctx)
	if err := p.WriteFieldBegin(ctx, name, thrift.STRING, field); err != nil {
		return thrift.PrependError("write field begin error: ", err)
	}
	if err := p.WriteBinary(ctx, value); err != nil {
		return thrift.PrependError("field write error: ", err)
	}
	if err := p.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError("write field end error: ", err)
	}
	return nil
}

// WriteStruct writes thrift.Struct of filed and id `name` and `field` respectively into `p`.
func WriteStruct(fctx FContext, p thrift.TProtocol, value thrift.TStruct, name string, field int16) error {
	ctx := toCTX(fctx)
	if err := p.WriteFieldBegin(ctx, name, thrift.STRUCT, field); err != nil {
		return thrift.PrependError("write field begin error: ", err)
	}
	if err := value.Write(ctx, p); err != nil {
		return thrift.PrependError("field write error: ", err)
	}
	if err := p.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError("write field end error: ", err)
	}
	return nil
}
