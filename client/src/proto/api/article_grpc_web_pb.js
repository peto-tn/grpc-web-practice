/**
 * @fileoverview gRPC-Web generated client stub for proto.api
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var proto_api_api_pb = require('../../proto/api/api_pb.js')

var proto_data_article_data_pb = require('../../proto/data/article_data_pb.js')
const proto = {};
proto.proto = {};
proto.proto.api = require('./article_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.api.ArticleClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.api.ArticlePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.api.GetArticleRequest,
 *   !proto.proto.api.GetArticleResponse>}
 */
const methodDescriptor_Article_Get = new grpc.web.MethodDescriptor(
  '/proto.api.Article/Get',
  grpc.web.MethodType.UNARY,
  proto.proto.api.GetArticleRequest,
  proto.proto.api.GetArticleResponse,
  /**
   * @param {!proto.proto.api.GetArticleRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.api.GetArticleResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.api.GetArticleRequest,
 *   !proto.proto.api.GetArticleResponse>}
 */
const methodInfo_Article_Get = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.api.GetArticleResponse,
  /**
   * @param {!proto.proto.api.GetArticleRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.api.GetArticleResponse.deserializeBinary
);


/**
 * @param {!proto.proto.api.GetArticleRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.api.GetArticleResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.api.GetArticleResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.api.ArticleClient.prototype.get =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.api.Article/Get',
      request,
      metadata || {},
      methodDescriptor_Article_Get,
      callback);
};


/**
 * @param {!proto.proto.api.GetArticleRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.api.GetArticleResponse>}
 *     A native promise that resolves to the response
 */
proto.proto.api.ArticlePromiseClient.prototype.get =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.api.Article/Get',
      request,
      metadata || {},
      methodDescriptor_Article_Get);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.api.ListArticleRequest,
 *   !proto.proto.api.ListArticleResponse>}
 */
const methodDescriptor_Article_List = new grpc.web.MethodDescriptor(
  '/proto.api.Article/List',
  grpc.web.MethodType.UNARY,
  proto.proto.api.ListArticleRequest,
  proto.proto.api.ListArticleResponse,
  /**
   * @param {!proto.proto.api.ListArticleRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.api.ListArticleResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.api.ListArticleRequest,
 *   !proto.proto.api.ListArticleResponse>}
 */
const methodInfo_Article_List = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.api.ListArticleResponse,
  /**
   * @param {!proto.proto.api.ListArticleRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.api.ListArticleResponse.deserializeBinary
);


/**
 * @param {!proto.proto.api.ListArticleRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.api.ListArticleResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.api.ListArticleResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.api.ArticleClient.prototype.list =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.api.Article/List',
      request,
      metadata || {},
      methodDescriptor_Article_List,
      callback);
};


/**
 * @param {!proto.proto.api.ListArticleRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.api.ListArticleResponse>}
 *     A native promise that resolves to the response
 */
proto.proto.api.ArticlePromiseClient.prototype.list =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.api.Article/List',
      request,
      metadata || {},
      methodDescriptor_Article_List);
};


module.exports = proto.proto.api;

