import 'package:flutter/material.dart';
import 'package:frontend/models/response_model.dart';
import 'package:frontend/models/user_model.dart';
import 'package:frontend/services/auth/auth_api_service.dart';


class AuthProvider extends ChangeNotifier {
  final AuthApiService _authService;
  bool _isLoading = false;
  Map<String ,dynamic>? _error;
  String? _message;
  LoginUserModel? _user;

  AuthProvider({required AuthApiService authService}) : _authService = authService;


  // Getters for the properties
  bool get isLoading => _isLoading;
  Map<String ,dynamic>? get error => _error;
  LoginUserModel? get user => _user;
  String? get message => _message;

  Future<bool> login(String email, String password) async {
    //  set the loading state to true 
    _isLoading = true;
    _error = null;
    notifyListeners();

    try {
      ResponseModel<LoginUserModel> result = await _authService.login(email, password);
      
      if (result.success) {
        _user = result.data;
        _error = null;
        _isLoading = false;
        _message = result.message;
        notifyListeners();
        return true;
      } else {
        _error =  result.errors;
        _isLoading = false;
        _message = result.message;
        notifyListeners();
        return false;
      }
    } catch (e) {
      _error = {'error': 'Login failed: $e'};
      _isLoading = false;
      notifyListeners();
      return false;
    }
  }

  void clearError() {
    _error = null;
    notifyListeners();
  }

  void logout() {
    _user = null;
    _error = null;
    notifyListeners();
  }
}