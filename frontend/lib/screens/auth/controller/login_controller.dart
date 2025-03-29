import 'package:flutter/material.dart';
import 'package:frontend/models/response_model.dart';
import 'package:frontend/models/user_model.dart';
import 'package:frontend/services/auth/auth_api_impl.dart';
// import 'package:frontend/services/auth/auth_api_service.dart';

class LoginController extends ChangeNotifier {
  final AuthApiImpl authService;

  LoginController({required this.authService});

  final TextEditingController emailController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();
  final GlobalKey<FormState> formKey = GlobalKey<FormState>();
  final Map<String, String?> errorText = {"email": null, "password": null};
  bool isPasswordVisible = false;
  bool isLoading = false;
  String? errorMessage;
  String? successMessage;

  void togglePasswordVisibility() {
    isPasswordVisible = !isPasswordVisible;
    notifyListeners();
  }

  void validateField(String field, String? errorMsg) {
    errorText[field] = errorMsg;
    notifyListeners();
  }
  
  Future<void> login() async {
    if (!formKey.currentState!.validate()) return;

    isLoading = true;
    errorMessage = null;
    successMessage = null;
    notifyListeners();

    try {
      ResponseModel<LoginUserModel> result = await authService.login(
        emailController.text,
        passwordController.text,
      );

      if (result.success) {
        successMessage = "Login successful!";
      } else {
        errorMessage = result.message ?? "Login failed!";
      }
    } catch (e) {
      errorMessage = "Login failed: $e";
    } finally {
      isLoading = false;
      notifyListeners();
    }
  }

  @override
  void dispose() {
    emailController.dispose();
    passwordController.dispose();
    super.dispose();
  }
}
