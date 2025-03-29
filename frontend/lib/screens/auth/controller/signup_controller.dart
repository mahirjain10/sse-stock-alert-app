import 'package:flutter/material.dart';
import 'package:frontend/models/response_model.dart';
import 'package:frontend/models/user_model.dart';
import 'package:frontend/services/auth/auth_api_impl.dart';

class SignupController extends ChangeNotifier {
    final AuthApiImpl authService;
  SignupController({required this.authService});

  final TextEditingController nameController = TextEditingController();
  final TextEditingController emailController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();
  final GlobalKey<FormState> formKey = GlobalKey<FormState>();
  final Map<String, String?> errorText = {
    "name": null,
    "email": null,
    "password": null
  };

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

  Future<void> signup() async {
    if (!formKey.currentState!.validate()) return;

    isLoading = true;
    errorMessage = null;
    successMessage = null;
    notifyListeners();
    try {
      ResponseModel<SignupUserModel> result = await authService.register(
        nameController.text,
        emailController.text,
        passwordController.text,
      );

      if (result.success) {
        successMessage = result.message ;
      } else {
        errorMessage = result.message ;
      }
    } catch (e) {
      errorMessage = "Sign up failed: $e";
    } finally {
      isLoading = false;
      notifyListeners();
    }
  }
}
