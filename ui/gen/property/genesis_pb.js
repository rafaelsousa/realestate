// source: property/genesis.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var property_owner_pb = require('../property/owner_pb.js');
goog.object.extend(proto, property_owner_pb);
var property_property_pb = require('../property/property_pb.js');
goog.object.extend(proto, property_property_pb);
goog.exportSymbol('proto.rafaelsousa.realestate.property.GenesisState', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.rafaelsousa.realestate.property.GenesisState = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.rafaelsousa.realestate.property.GenesisState.repeatedFields_, null);
};
goog.inherits(proto.rafaelsousa.realestate.property.GenesisState, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.rafaelsousa.realestate.property.GenesisState.displayName = 'proto.rafaelsousa.realestate.property.GenesisState';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.rafaelsousa.realestate.property.GenesisState.repeatedFields_ = [2,1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.rafaelsousa.realestate.property.GenesisState.prototype.toObject = function(opt_includeInstance) {
  return proto.rafaelsousa.realestate.property.GenesisState.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.rafaelsousa.realestate.property.GenesisState} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.rafaelsousa.realestate.property.GenesisState.toObject = function(includeInstance, msg) {
  var f, obj = {
    ownerlistList: jspb.Message.toObjectList(msg.getOwnerlistList(),
    property_owner_pb.Owner.toObject, includeInstance),
    propertylistList: jspb.Message.toObjectList(msg.getPropertylistList(),
    property_property_pb.Property.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.rafaelsousa.realestate.property.GenesisState}
 */
proto.rafaelsousa.realestate.property.GenesisState.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.rafaelsousa.realestate.property.GenesisState;
  return proto.rafaelsousa.realestate.property.GenesisState.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.rafaelsousa.realestate.property.GenesisState} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.rafaelsousa.realestate.property.GenesisState}
 */
proto.rafaelsousa.realestate.property.GenesisState.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 2:
      var value = new property_owner_pb.Owner;
      reader.readMessage(value,property_owner_pb.Owner.deserializeBinaryFromReader);
      msg.addOwnerlist(value);
      break;
    case 1:
      var value = new property_property_pb.Property;
      reader.readMessage(value,property_property_pb.Property.deserializeBinaryFromReader);
      msg.addPropertylist(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.rafaelsousa.realestate.property.GenesisState.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.rafaelsousa.realestate.property.GenesisState.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.rafaelsousa.realestate.property.GenesisState} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.rafaelsousa.realestate.property.GenesisState.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getOwnerlistList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      property_owner_pb.Owner.serializeBinaryToWriter
    );
  }
  f = message.getPropertylistList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      property_property_pb.Property.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Owner ownerList = 2;
 * @return {!Array<!proto.rafaelsousa.realestate.property.Owner>}
 */
proto.rafaelsousa.realestate.property.GenesisState.prototype.getOwnerlistList = function() {
  return /** @type{!Array<!proto.rafaelsousa.realestate.property.Owner>} */ (
    jspb.Message.getRepeatedWrapperField(this, property_owner_pb.Owner, 2));
};


/**
 * @param {!Array<!proto.rafaelsousa.realestate.property.Owner>} value
 * @return {!proto.rafaelsousa.realestate.property.GenesisState} returns this
*/
proto.rafaelsousa.realestate.property.GenesisState.prototype.setOwnerlistList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.rafaelsousa.realestate.property.Owner=} opt_value
 * @param {number=} opt_index
 * @return {!proto.rafaelsousa.realestate.property.Owner}
 */
proto.rafaelsousa.realestate.property.GenesisState.prototype.addOwnerlist = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.rafaelsousa.realestate.property.Owner, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.rafaelsousa.realestate.property.GenesisState} returns this
 */
proto.rafaelsousa.realestate.property.GenesisState.prototype.clearOwnerlistList = function() {
  return this.setOwnerlistList([]);
};


/**
 * repeated Property propertyList = 1;
 * @return {!Array<!proto.rafaelsousa.realestate.property.Property>}
 */
proto.rafaelsousa.realestate.property.GenesisState.prototype.getPropertylistList = function() {
  return /** @type{!Array<!proto.rafaelsousa.realestate.property.Property>} */ (
    jspb.Message.getRepeatedWrapperField(this, property_property_pb.Property, 1));
};


/**
 * @param {!Array<!proto.rafaelsousa.realestate.property.Property>} value
 * @return {!proto.rafaelsousa.realestate.property.GenesisState} returns this
*/
proto.rafaelsousa.realestate.property.GenesisState.prototype.setPropertylistList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.rafaelsousa.realestate.property.Property=} opt_value
 * @param {number=} opt_index
 * @return {!proto.rafaelsousa.realestate.property.Property}
 */
proto.rafaelsousa.realestate.property.GenesisState.prototype.addPropertylist = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.rafaelsousa.realestate.property.Property, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.rafaelsousa.realestate.property.GenesisState} returns this
 */
proto.rafaelsousa.realestate.property.GenesisState.prototype.clearPropertylistList = function() {
  return this.setPropertylistList([]);
};


goog.object.extend(exports, proto.rafaelsousa.realestate.property);
