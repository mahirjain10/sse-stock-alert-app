import 'package:frontend/models/response_model.dart';
import 'package:frontend/models/user_model.dart';

abstract class AuthApiService {
  Future<ResponseModel<LoginUserModel>> login(String email, String password);
  Future<ResponseModel<SignupUserModel>> register(String name,String email, String password);
}
