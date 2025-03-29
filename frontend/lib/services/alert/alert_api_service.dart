import 'package:frontend/models/create_alert.model.dart';
import 'package:frontend/models/response_model.dart';

abstract class AlertApiService {
  Future<ResponseModel<CreateAlertModel>> createAlert(CreateAlertModel alert);
  Future<ResponseModel<GetCurrentPriceAndTimeModel>> getcurrentPriceAndTime(GetCurrentPriceAndTimeModel getCPT);
  // Future<ResponseModel<List<CreateAlertModel>>> getAlerts(String userId);
  // Future<ResponseModel<CreateAlertModel>> getAlert(String userId, String alertId);
  // Future<ResponseModel<void>> deleteAlert(String userId, String alertId);
}