import 'dart:convert';
import 'package:dio/dio.dart';
import 'package:cookie_jar/cookie_jar.dart';
import 'package:dio_cookie_manager/dio_cookie_manager.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter/foundation.dart';
import 'package:frontend/models/create_alert.model.dart';
import 'package:frontend/models/response_model.dart';
import 'package:frontend/services/alert/alert_api_service.dart';

class AlertApiImpl implements AlertApiService {
  final Dio dio;
  final CookieJar cookieJar = CookieJar();

  AlertApiImpl(this.dio) {
    dio.options.baseUrl = dotenv.env["BASE_URL"]!;
    dio.options.headers["Content-Type"] = "application/json";
    dio.interceptors
        .add(CookieManager(cookieJar)); // Automatically handle cookies
  }

  @override
  Future<ResponseModel<CreateAlertModel>> createAlert(
      CreateAlertModel alert) async {
    try {
      String requestBody = alert.toJsonString();

      final response = await dio.post(
        "/api/alert/create-stock-alert",
        data: requestBody,
      );

      print("Response status code: ${response.statusCode}");
      print("Response body: ${response.data}");

      return ResponseModel<CreateAlertModel>.fromJson(
        jsonEncode(response.data),
        response.statusCode!,
        response.data,
        (data) => CreateAlertModel.fromJson(data),
      );
    } catch (e) {
      throw Exception("Error during create alert: $e");
    }
  }

  @override
  Future<ResponseModel<GetCurrentPriceAndTimeModel>> getcurrentPriceAndTime(
      GetCurrentPriceAndTimeModel getCPT) async {
    try {
      String requestBody = getCPT.toJsonString();

      final response = await dio.post(
        "/api/alert/get-current-price",
        data: requestBody,
      );

      debugPrint("Response status code: ${response.statusCode}");
      debugPrint("Response body: ${response.data}");

      return ResponseModel<GetCurrentPriceAndTimeModel>.fromJson(
        jsonEncode(response.data),
        response.statusCode!,
        response.data,
        (data) => GetCurrentPriceAndTimeModel.fromJson(data),
      );
    } catch (e) {
      throw Exception("Error during get current price: $e");
    }
  }
}
