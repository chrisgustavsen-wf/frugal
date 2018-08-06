/**
 * Autogenerated by Frugal Compiler (2.21.0)
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *
 * @generated
 */
package v1.music;

import org.apache.thrift.scheme.IScheme;
import org.apache.thrift.scheme.SchemeFactory;
import org.apache.thrift.scheme.StandardScheme;

import org.apache.thrift.scheme.TupleScheme;
import org.apache.thrift.protocol.TTupleProtocol;
import org.apache.thrift.protocol.TProtocolException;
import org.apache.thrift.EncodingUtils;
import org.apache.thrift.TException;
import org.apache.thrift.async.AsyncMethodCallback;
import org.apache.thrift.server.AbstractNonblockingServer.*;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;
import java.util.EnumMap;
import java.util.Set;
import java.util.HashSet;
import java.util.EnumSet;
import java.util.Collections;
import java.util.BitSet;
import java.nio.ByteBuffer;
import java.util.Arrays;
import javax.annotation.Generated;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * Comments (with an @ symbol) will be added to generated code.
 */
@Generated(value = "Autogenerated by Frugal Compiler (2.21.0)")
public class Track implements org.apache.thrift.TBase<Track, Track._Fields>, java.io.Serializable, Cloneable, Comparable<Track> {
	private static final org.apache.thrift.protocol.TStruct STRUCT_DESC = new org.apache.thrift.protocol.TStruct("Track");

	private static final org.apache.thrift.protocol.TField TITLE_FIELD_DESC = new org.apache.thrift.protocol.TField("title", org.apache.thrift.protocol.TType.STRING, (short)1);
	private static final org.apache.thrift.protocol.TField ARTIST_FIELD_DESC = new org.apache.thrift.protocol.TField("artist", org.apache.thrift.protocol.TType.STRING, (short)2);
	private static final org.apache.thrift.protocol.TField PUBLISHER_FIELD_DESC = new org.apache.thrift.protocol.TField("publisher", org.apache.thrift.protocol.TType.STRING, (short)3);
	private static final org.apache.thrift.protocol.TField COMPOSER_FIELD_DESC = new org.apache.thrift.protocol.TField("composer", org.apache.thrift.protocol.TType.STRING, (short)4);
	private static final org.apache.thrift.protocol.TField DURATION_FIELD_DESC = new org.apache.thrift.protocol.TField("duration", org.apache.thrift.protocol.TType.DOUBLE, (short)5);
	private static final org.apache.thrift.protocol.TField PRO_FIELD_DESC = new org.apache.thrift.protocol.TField("pro", org.apache.thrift.protocol.TType.I32, (short)6);

	private static final Map<Class<? extends IScheme>, SchemeFactory> schemes = new HashMap<Class<? extends IScheme>, SchemeFactory>();
	static {
		schemes.put(StandardScheme.class, new TrackStandardSchemeFactory());
		schemes.put(TupleScheme.class, new TrackTupleSchemeFactory());
	}

	public String title;
	public String artist;
	public String publisher;
	public String composer;
	public double duration;
	public PerfRightsOrg pro;
	/** The set of fields this struct contains, along with convenience methods for finding and manipulating them. */
	public enum _Fields implements org.apache.thrift.TFieldIdEnum {
		TITLE((short)1, "title"),
		ARTIST((short)2, "artist"),
		PUBLISHER((short)3, "publisher"),
		COMPOSER((short)4, "composer"),
		DURATION((short)5, "duration"),
		PRO((short)6, "pro")
		;

		private static final Map<String, _Fields> byName = new HashMap<String, _Fields>();

		static {
			for (_Fields field : EnumSet.allOf(_Fields.class)) {
				byName.put(field.getFieldName(), field);
			}
		}

		/**
		 * Find the _Fields constant that matches fieldId, or null if its not found.
		 */
		public static _Fields findByThriftId(int fieldId) {
			switch(fieldId) {
				case 1: // TITLE
					return TITLE;
				case 2: // ARTIST
					return ARTIST;
				case 3: // PUBLISHER
					return PUBLISHER;
				case 4: // COMPOSER
					return COMPOSER;
				case 5: // DURATION
					return DURATION;
				case 6: // PRO
					return PRO;
				default:
					return null;
			}
		}

		/**
		 * Find the _Fields constant that matches fieldId, throwing an exception
		 * if it is not found.
		 */
		public static _Fields findByThriftIdOrThrow(int fieldId) {
			_Fields fields = findByThriftId(fieldId);
			if (fields == null) throw new IllegalArgumentException("Field " + fieldId + " doesn't exist!");
			return fields;
		}

		/**
		 * Find the _Fields constant that matches name, or null if its not found.
		 */
		public static _Fields findByName(String name) {
			return byName.get(name);
		}

		private final short _thriftId;
		private final String _fieldName;

		_Fields(short thriftId, String fieldName) {
			_thriftId = thriftId;
			_fieldName = fieldName;
		}

		public short getThriftFieldId() {
			return _thriftId;
		}

		public String getFieldName() {
			return _fieldName;
		}
	}

	// isset id assignments
	private static final int __DURATION_ISSET_ID = 0;
	private byte __isset_bitfield = 0;
	public Track() {
	}

	public Track(
		String title,
		String artist,
		String publisher,
		String composer,
		double duration,
		PerfRightsOrg pro) {
		this();
		this.title = title;
		this.artist = artist;
		this.publisher = publisher;
		this.composer = composer;
		this.duration = duration;
		setDurationIsSet(true);
		this.pro = pro;
	}

	/**
	 * Performs a deep copy on <i>other</i>.
	 */
	public Track(Track other) {
		__isset_bitfield = other.__isset_bitfield;
		if (other.isSetTitle()) {
			this.title = other.title;
		}
		if (other.isSetArtist()) {
			this.artist = other.artist;
		}
		if (other.isSetPublisher()) {
			this.publisher = other.publisher;
		}
		if (other.isSetComposer()) {
			this.composer = other.composer;
		}
		this.duration = other.duration;
		if (other.isSetPro()) {
			this.pro = other.pro;
		}
	}

	public Track deepCopy() {
		return new Track(this);
	}

	@Override
	public void clear() {
		this.title = null;

		this.artist = null;

		this.publisher = null;

		this.composer = null;

		setDurationIsSet(false);
		this.duration = 0.0;

		this.pro = null;

	}

	public String getTitle() {
		return this.title;
	}

	public Track setTitle(String title) {
		this.title = title;
		return this;
	}

	public void unsetTitle() {
		this.title = null;
	}

	/** Returns true if field title is set (has been assigned a value) and false otherwise */
	public boolean isSetTitle() {
		return this.title != null;
	}

	public void setTitleIsSet(boolean value) {
		if (!value) {
			this.title = null;
		}
	}

	public String getArtist() {
		return this.artist;
	}

	public Track setArtist(String artist) {
		this.artist = artist;
		return this;
	}

	public void unsetArtist() {
		this.artist = null;
	}

	/** Returns true if field artist is set (has been assigned a value) and false otherwise */
	public boolean isSetArtist() {
		return this.artist != null;
	}

	public void setArtistIsSet(boolean value) {
		if (!value) {
			this.artist = null;
		}
	}

	public String getPublisher() {
		return this.publisher;
	}

	public Track setPublisher(String publisher) {
		this.publisher = publisher;
		return this;
	}

	public void unsetPublisher() {
		this.publisher = null;
	}

	/** Returns true if field publisher is set (has been assigned a value) and false otherwise */
	public boolean isSetPublisher() {
		return this.publisher != null;
	}

	public void setPublisherIsSet(boolean value) {
		if (!value) {
			this.publisher = null;
		}
	}

	public String getComposer() {
		return this.composer;
	}

	public Track setComposer(String composer) {
		this.composer = composer;
		return this;
	}

	public void unsetComposer() {
		this.composer = null;
	}

	/** Returns true if field composer is set (has been assigned a value) and false otherwise */
	public boolean isSetComposer() {
		return this.composer != null;
	}

	public void setComposerIsSet(boolean value) {
		if (!value) {
			this.composer = null;
		}
	}

	public double getDuration() {
		return this.duration;
	}

	public Track setDuration(double duration) {
		this.duration = duration;
		setDurationIsSet(true);
		return this;
	}

	public void unsetDuration() {
		__isset_bitfield = EncodingUtils.clearBit(__isset_bitfield, __DURATION_ISSET_ID);
	}

	/** Returns true if field duration is set (has been assigned a value) and false otherwise */
	public boolean isSetDuration() {
		return EncodingUtils.testBit(__isset_bitfield, __DURATION_ISSET_ID);
	}

	public void setDurationIsSet(boolean value) {
		__isset_bitfield = EncodingUtils.setBit(__isset_bitfield, __DURATION_ISSET_ID, value);
	}

	public PerfRightsOrg getPro() {
		return this.pro;
	}

	public Track setPro(PerfRightsOrg pro) {
		this.pro = pro;
		return this;
	}

	public void unsetPro() {
		this.pro = null;
	}

	/** Returns true if field pro is set (has been assigned a value) and false otherwise */
	public boolean isSetPro() {
		return this.pro != null;
	}

	public void setProIsSet(boolean value) {
		if (!value) {
			this.pro = null;
		}
	}

	public void setFieldValue(_Fields field, Object value) {
		switch (field) {
		case TITLE:
			if (value == null) {
				unsetTitle();
			} else {
				setTitle((String)value);
			}
			break;

		case ARTIST:
			if (value == null) {
				unsetArtist();
			} else {
				setArtist((String)value);
			}
			break;

		case PUBLISHER:
			if (value == null) {
				unsetPublisher();
			} else {
				setPublisher((String)value);
			}
			break;

		case COMPOSER:
			if (value == null) {
				unsetComposer();
			} else {
				setComposer((String)value);
			}
			break;

		case DURATION:
			if (value == null) {
				unsetDuration();
			} else {
				setDuration((Double)value);
			}
			break;

		case PRO:
			if (value == null) {
				unsetPro();
			} else {
				setPro((PerfRightsOrg)value);
			}
			break;

		}
	}

	public Object getFieldValue(_Fields field) {
		switch (field) {
		case TITLE:
			return getTitle();

		case ARTIST:
			return getArtist();

		case PUBLISHER:
			return getPublisher();

		case COMPOSER:
			return getComposer();

		case DURATION:
			return getDuration();

		case PRO:
			return getPro();

		}
		throw new IllegalStateException();
	}

	/** Returns true if field corresponding to fieldID is set (has been assigned a value) and false otherwise */
	public boolean isSet(_Fields field) {
		if (field == null) {
			throw new IllegalArgumentException();
		}

		switch (field) {
		case TITLE:
			return isSetTitle();
		case ARTIST:
			return isSetArtist();
		case PUBLISHER:
			return isSetPublisher();
		case COMPOSER:
			return isSetComposer();
		case DURATION:
			return isSetDuration();
		case PRO:
			return isSetPro();
		}
		throw new IllegalStateException();
	}

	@Override
	public boolean equals(Object that) {
		if (that == null)
			return false;
		if (that instanceof Track)
			return this.equals((Track)that);
		return false;
	}

	public boolean equals(Track that) {
		if (that == null)
			return false;

		boolean this_present_title = true && this.isSetTitle();
		boolean that_present_title = true && that.isSetTitle();
		if (this_present_title || that_present_title) {
			if (!(this_present_title && that_present_title))
				return false;
			if (!this.title.equals(that.title))
				return false;
		}

		boolean this_present_artist = true && this.isSetArtist();
		boolean that_present_artist = true && that.isSetArtist();
		if (this_present_artist || that_present_artist) {
			if (!(this_present_artist && that_present_artist))
				return false;
			if (!this.artist.equals(that.artist))
				return false;
		}

		boolean this_present_publisher = true && this.isSetPublisher();
		boolean that_present_publisher = true && that.isSetPublisher();
		if (this_present_publisher || that_present_publisher) {
			if (!(this_present_publisher && that_present_publisher))
				return false;
			if (!this.publisher.equals(that.publisher))
				return false;
		}

		boolean this_present_composer = true && this.isSetComposer();
		boolean that_present_composer = true && that.isSetComposer();
		if (this_present_composer || that_present_composer) {
			if (!(this_present_composer && that_present_composer))
				return false;
			if (!this.composer.equals(that.composer))
				return false;
		}

		boolean this_present_duration = true;
		boolean that_present_duration = true;
		if (this_present_duration || that_present_duration) {
			if (!(this_present_duration && that_present_duration))
				return false;
			if (this.duration != that.duration)
				return false;
		}

		boolean this_present_pro = true && this.isSetPro();
		boolean that_present_pro = true && that.isSetPro();
		if (this_present_pro || that_present_pro) {
			if (!(this_present_pro && that_present_pro))
				return false;
			if (!this.pro.equals(that.pro))
				return false;
		}

		return true;
	}

	@Override
	public int hashCode() {
		List<Object> list = new ArrayList<Object>();

		boolean present_title = true && (isSetTitle());
		list.add(present_title);
		if (present_title)
			list.add(title);

		boolean present_artist = true && (isSetArtist());
		list.add(present_artist);
		if (present_artist)
			list.add(artist);

		boolean present_publisher = true && (isSetPublisher());
		list.add(present_publisher);
		if (present_publisher)
			list.add(publisher);

		boolean present_composer = true && (isSetComposer());
		list.add(present_composer);
		if (present_composer)
			list.add(composer);

		boolean present_duration = true;
		list.add(present_duration);
		if (present_duration)
			list.add(duration);

		boolean present_pro = true && (isSetPro());
		list.add(present_pro);
		if (present_pro)
			list.add(pro.getValue());

		return list.hashCode();
	}

	@Override
	public int compareTo(Track other) {
		if (!getClass().equals(other.getClass())) {
			return getClass().getName().compareTo(other.getClass().getName());
		}

		int lastComparison = 0;

		lastComparison = Boolean.valueOf(isSetTitle()).compareTo(other.isSetTitle());
		if (lastComparison != 0) {
			return lastComparison;
		}
		if (isSetTitle()) {
			lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.title, other.title);
			if (lastComparison != 0) {
				return lastComparison;
			}
		}
		lastComparison = Boolean.valueOf(isSetArtist()).compareTo(other.isSetArtist());
		if (lastComparison != 0) {
			return lastComparison;
		}
		if (isSetArtist()) {
			lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.artist, other.artist);
			if (lastComparison != 0) {
				return lastComparison;
			}
		}
		lastComparison = Boolean.valueOf(isSetPublisher()).compareTo(other.isSetPublisher());
		if (lastComparison != 0) {
			return lastComparison;
		}
		if (isSetPublisher()) {
			lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.publisher, other.publisher);
			if (lastComparison != 0) {
				return lastComparison;
			}
		}
		lastComparison = Boolean.valueOf(isSetComposer()).compareTo(other.isSetComposer());
		if (lastComparison != 0) {
			return lastComparison;
		}
		if (isSetComposer()) {
			lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.composer, other.composer);
			if (lastComparison != 0) {
				return lastComparison;
			}
		}
		lastComparison = Boolean.valueOf(isSetDuration()).compareTo(other.isSetDuration());
		if (lastComparison != 0) {
			return lastComparison;
		}
		if (isSetDuration()) {
			lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.duration, other.duration);
			if (lastComparison != 0) {
				return lastComparison;
			}
		}
		lastComparison = Boolean.valueOf(isSetPro()).compareTo(other.isSetPro());
		if (lastComparison != 0) {
			return lastComparison;
		}
		if (isSetPro()) {
			lastComparison = org.apache.thrift.TBaseHelper.compareTo(this.pro, other.pro);
			if (lastComparison != 0) {
				return lastComparison;
			}
		}
		return 0;
	}

	public _Fields fieldForId(int fieldId) {
		return _Fields.findByThriftId(fieldId);
	}

	public void read(org.apache.thrift.protocol.TProtocol iprot) throws org.apache.thrift.TException {
		schemes.get(iprot.getScheme()).getScheme().read(iprot, this);
	}

	public void write(org.apache.thrift.protocol.TProtocol oprot) throws org.apache.thrift.TException {
		schemes.get(oprot.getScheme()).getScheme().write(oprot, this);
	}

	@Override
	public String toString() {
		StringBuilder sb = new StringBuilder("Track(");
		boolean first = true;

		sb.append("title:");
		if (this.title == null) {
			sb.append("null");
		} else {
			sb.append(this.title);
		}
		first = false;
		if (!first) sb.append(", ");
		sb.append("artist:");
		if (this.artist == null) {
			sb.append("null");
		} else {
			sb.append(this.artist);
		}
		first = false;
		if (!first) sb.append(", ");
		sb.append("publisher:");
		if (this.publisher == null) {
			sb.append("null");
		} else {
			sb.append(this.publisher);
		}
		first = false;
		if (!first) sb.append(", ");
		sb.append("composer:");
		if (this.composer == null) {
			sb.append("null");
		} else {
			sb.append(this.composer);
		}
		first = false;
		if (!first) sb.append(", ");
		sb.append("duration:");
		sb.append(this.duration);
		first = false;
		if (!first) sb.append(", ");
		sb.append("pro:");
		if (this.pro == null) {
			sb.append("null");
		} else {
			sb.append(this.pro);
		}
		first = false;
		sb.append(")");
		return sb.toString();
	}

	public void validate() throws org.apache.thrift.TException {
		// check for required fields
		// check for sub-struct validity
	}

	private void writeObject(java.io.ObjectOutputStream out) throws java.io.IOException {
		try {
			write(new org.apache.thrift.protocol.TCompactProtocol(new org.apache.thrift.transport.TIOStreamTransport(out)));
		} catch (org.apache.thrift.TException te) {
			throw new java.io.IOException(te);
		}
	}

	private void readObject(java.io.ObjectInputStream in) throws java.io.IOException, ClassNotFoundException {
		try {
			// it doesn't seem like you should have to do this, but java serialization is wacky, and doesn't call the default constructor.
			__isset_bitfield = 0;
			read(new org.apache.thrift.protocol.TCompactProtocol(new org.apache.thrift.transport.TIOStreamTransport(in)));
		} catch (org.apache.thrift.TException te) {
			throw new java.io.IOException(te);
		}
	}

	private static class TrackStandardSchemeFactory implements SchemeFactory {
		public TrackStandardScheme getScheme() {
			return new TrackStandardScheme();
		}
	}

	private static class TrackStandardScheme extends StandardScheme<Track> {

		public void read(org.apache.thrift.protocol.TProtocol iprot, Track struct) throws org.apache.thrift.TException {
			org.apache.thrift.protocol.TField schemeField;
			iprot.readStructBegin();
			while (true) {
				schemeField = iprot.readFieldBegin();
				if (schemeField.type == org.apache.thrift.protocol.TType.STOP) {
					break;
				}
				switch (schemeField.id) {
					case 1: // TITLE
						if (schemeField.type == org.apache.thrift.protocol.TType.STRING) {
							struct.title = iprot.readString();
							struct.setTitleIsSet(true);
						} else {
							org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
						}
						break;
					case 2: // ARTIST
						if (schemeField.type == org.apache.thrift.protocol.TType.STRING) {
							struct.artist = iprot.readString();
							struct.setArtistIsSet(true);
						} else {
							org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
						}
						break;
					case 3: // PUBLISHER
						if (schemeField.type == org.apache.thrift.protocol.TType.STRING) {
							struct.publisher = iprot.readString();
							struct.setPublisherIsSet(true);
						} else {
							org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
						}
						break;
					case 4: // COMPOSER
						if (schemeField.type == org.apache.thrift.protocol.TType.STRING) {
							struct.composer = iprot.readString();
							struct.setComposerIsSet(true);
						} else {
							org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
						}
						break;
					case 5: // DURATION
						if (schemeField.type == org.apache.thrift.protocol.TType.DOUBLE) {
							struct.duration = iprot.readDouble();
							struct.setDurationIsSet(true);
						} else {
							org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
						}
						break;
					case 6: // PRO
						if (schemeField.type == org.apache.thrift.protocol.TType.I32) {
							struct.pro = PerfRightsOrg.findByValue(iprot.readI32());
							struct.setProIsSet(true);
						} else {
							org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
						}
						break;
					default:
						org.apache.thrift.protocol.TProtocolUtil.skip(iprot, schemeField.type);
				}
				iprot.readFieldEnd();
			}
			iprot.readStructEnd();

			// check for required fields of primitive type, which can't be checked in the validate method
			struct.validate();
		}

		public void write(org.apache.thrift.protocol.TProtocol oprot, Track struct) throws org.apache.thrift.TException {
			struct.validate();

			oprot.writeStructBegin(STRUCT_DESC);
			if (struct.title != null) {
				oprot.writeFieldBegin(TITLE_FIELD_DESC);
				String elem0 = struct.title;
				oprot.writeString(elem0);
				oprot.writeFieldEnd();
			}
			if (struct.artist != null) {
				oprot.writeFieldBegin(ARTIST_FIELD_DESC);
				String elem1 = struct.artist;
				oprot.writeString(elem1);
				oprot.writeFieldEnd();
			}
			if (struct.publisher != null) {
				oprot.writeFieldBegin(PUBLISHER_FIELD_DESC);
				String elem2 = struct.publisher;
				oprot.writeString(elem2);
				oprot.writeFieldEnd();
			}
			if (struct.composer != null) {
				oprot.writeFieldBegin(COMPOSER_FIELD_DESC);
				String elem3 = struct.composer;
				oprot.writeString(elem3);
				oprot.writeFieldEnd();
			}
			oprot.writeFieldBegin(DURATION_FIELD_DESC);
			double elem4 = struct.duration;
			oprot.writeDouble(elem4);
			oprot.writeFieldEnd();
			if (struct.pro != null) {
				oprot.writeFieldBegin(PRO_FIELD_DESC);
				PerfRightsOrg elem5 = struct.pro;
				oprot.writeI32(elem5.getValue());
				oprot.writeFieldEnd();
			}
			oprot.writeFieldStop();
			oprot.writeStructEnd();
		}

	}

	private static class TrackTupleSchemeFactory implements SchemeFactory {
		public TrackTupleScheme getScheme() {
			return new TrackTupleScheme();
		}
	}

	private static class TrackTupleScheme extends TupleScheme<Track> {

		@Override
		public void write(org.apache.thrift.protocol.TProtocol prot, Track struct) throws org.apache.thrift.TException {
			TTupleProtocol oprot = (TTupleProtocol) prot;
			BitSet optionals = new BitSet();
			if (struct.isSetTitle()) {
				optionals.set(0);
			}
			if (struct.isSetArtist()) {
				optionals.set(1);
			}
			if (struct.isSetPublisher()) {
				optionals.set(2);
			}
			if (struct.isSetComposer()) {
				optionals.set(3);
			}
			if (struct.isSetDuration()) {
				optionals.set(4);
			}
			if (struct.isSetPro()) {
				optionals.set(5);
			}
			oprot.writeBitSet(optionals, 6);
			if (struct.isSetTitle()) {
				String elem6 = struct.title;
				oprot.writeString(elem6);
			}
			if (struct.isSetArtist()) {
				String elem7 = struct.artist;
				oprot.writeString(elem7);
			}
			if (struct.isSetPublisher()) {
				String elem8 = struct.publisher;
				oprot.writeString(elem8);
			}
			if (struct.isSetComposer()) {
				String elem9 = struct.composer;
				oprot.writeString(elem9);
			}
			if (struct.isSetDuration()) {
				double elem10 = struct.duration;
				oprot.writeDouble(elem10);
			}
			if (struct.isSetPro()) {
				PerfRightsOrg elem11 = struct.pro;
				oprot.writeI32(elem11.getValue());
			}
		}

		@Override
		public void read(org.apache.thrift.protocol.TProtocol prot, Track struct) throws org.apache.thrift.TException {
			TTupleProtocol iprot = (TTupleProtocol) prot;
			BitSet incoming = iprot.readBitSet(6);
			if (incoming.get(0)) {
				struct.title = iprot.readString();
				struct.setTitleIsSet(true);
			}
			if (incoming.get(1)) {
				struct.artist = iprot.readString();
				struct.setArtistIsSet(true);
			}
			if (incoming.get(2)) {
				struct.publisher = iprot.readString();
				struct.setPublisherIsSet(true);
			}
			if (incoming.get(3)) {
				struct.composer = iprot.readString();
				struct.setComposerIsSet(true);
			}
			if (incoming.get(4)) {
				struct.duration = iprot.readDouble();
				struct.setDurationIsSet(true);
			}
			if (incoming.get(5)) {
				struct.pro = PerfRightsOrg.findByValue(iprot.readI32());
				struct.setProIsSet(true);
			}
		}

	}

}
