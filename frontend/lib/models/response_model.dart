import 'package:flutter/foundation.dart';

class ResponseModel<T> {
  final int statusCode;
  final String responseJson;
  final T? data;
  final Map<String, String>? errors;
  final String message;
  final bool success;

  ResponseModel({
    required this.statusCode,
    required this.responseJson,
    this.data,
    this.errors,
    required this.message,
    required this.success,
  });

  factory ResponseModel.fromJson(
    String responseJson, // takes the whole JSON response
    int statusCode, // takes the status code
    Map<String, dynamic> json, // takes the JSON response of data
    T Function(Map<String, dynamic>)?
        fromJsonT, // takes the custom function on how to deseralize the data.
  ) {

    
    // Get success value directly from the API response
    final bool success = json['success'] as bool? ?? false;
    final String message = json['message'] as String;

    
    // Parse data if present and a fromJsonT function is provided
    T? parsedData;
    if (json.containsKey('data') && fromJsonT != null && json['data'] != null) {
      debugPrint("json['data'] ${json['data']}");
      if (json['data'] is Map<String, dynamic>) {
        parsedData = fromJsonT(json['data'] as Map<String, dynamic>);
      }
    }

    // Parse errors if present
    Map<String, String>? errors;
    if (json.containsKey('errors') && json['errors'] != null) {
      if (json['errors'] is Map<String, String>) {
        errors = json['errors'] as Map<String, String>;
      }
    }

    return ResponseModel<T>(
      statusCode: statusCode,
      responseJson: responseJson,
      data: parsedData,
      errors: errors,
      success: success,
      message: message,
    );
  }
}
