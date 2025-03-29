import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:frontend/main.dart';
import 'package:frontend/screens/auth/presentation/forget_password_text.dart';
import 'package:frontend/screens/auth/presentation/signup_login_text.dart';
import 'package:frontend/screens/auth/presentation/welcome_text.dart';
// import 'package:frontend/screens/create_alert/create_alert_page.dart';
import 'package:frontend/services/auth/auth_api_impl.dart';
import 'package:frontend/widgets/auth_button.dart';
import 'package:frontend/widgets/auth_input_bar.dart';
import 'package:frontend/widgets/auth_input_label.dart';
import 'package:provider/provider.dart';
import 'package:toastification/toastification.dart';
import 'package:frontend/screens/auth/controller/login_controller.dart';

class LoginPage extends StatefulWidget {
  final Dio dio;

  const LoginPage({super.key, required this.dio});

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  void _showToast(BuildContext context, String message, bool isError) {
    debugPrint(message);
    toastification.show(
      context: context,
      type: isError ? ToastificationType.error : ToastificationType.success,
      style: ToastificationStyle.flat,
      title: Text(isError ? "Error" : "Success",
          style: const TextStyle(fontWeight: FontWeight.bold)),
      description: Text(message),
      alignment: Alignment.topCenter,
      autoCloseDuration: const Duration(seconds: 4),
    );
  }

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (context) => LoginController(authService: AuthApiImpl(dio)),
      child: Consumer<LoginController>(
        builder: (context, loginController, child) {
          return SafeArea(
            child: Scaffold(
              backgroundColor: const Color(0xfff3f4f6),
              body: SingleChildScrollView(
                child: Center(
                  child: Column(
                    children: [
                      Padding(
                        padding: EdgeInsets.fromLTRB(
                          0,
                          MediaQuery.sizeOf(context).height * 0.2,
                          0,
                          MediaQuery.sizeOf(context).height * 0.05,
                        ),
                        child: const WelcomeText(),
                      ),
                      SizedBox(
                        height: MediaQuery.sizeOf(context).height * 0.5,
                        child: Form(
                          key: loginController.formKey,
                          child: Column(
                            children: [
                              SizedBox(
                                height:
                                    MediaQuery.sizeOf(context).height * 0.25,
                                child: Column(
                                  crossAxisAlignment: CrossAxisAlignment.center,
                                  mainAxisAlignment:
                                      MainAxisAlignment.spaceEvenly,
                                  children: [
                                    Column(
                                      crossAxisAlignment:
                                          CrossAxisAlignment.start,
                                      children: [
                                        const Padding(
                                          padding: EdgeInsets.only(bottom: 5),
                                          child: AuthInputLabel(label: "Email"),
                                        ),
                                        InputField(
                                          hintText: "Enter your email",
                                          field: "email",
                                          isVisible: false,
                                          validateForm: (field, errorMsg) =>
                                              loginController.validateField(
                                                  field, errorMsg),
                                          onToggle: () {},
                                          controller:
                                              loginController.emailController,
                                          errorText: loginController.errorText,
                                        ),
                                      ],
                                    ),
                                    Column(
                                      crossAxisAlignment:
                                          CrossAxisAlignment.start,
                                      children: [
                                        const Padding(
                                          padding: EdgeInsets.only(bottom: 5),
                                          child:
                                              AuthInputLabel(label: "Password"),
                                        ),
                                        InputField(
                                          hintText: "Enter your Password",
                                          field: "password",
                                          isVisible:
                                              loginController.isPasswordVisible,
                                          validateForm: (field, errorMsg) =>
                                              loginController.validateField(
                                                  field, errorMsg),
                                          onToggle: loginController
                                              .togglePasswordVisibility,
                                          controller: loginController
                                              .passwordController,
                                          errorText: loginController.errorText,
                                        ),
                                      ],
                                    ),
                                    Align(
                                      alignment: Alignment.centerRight,
                                      child: Padding(
                                        padding:
                                            const EdgeInsets.only(right: 25),
                                        child: ForgetPassword(onTap: () => {}),
                                      ),
                                    ),
                                  ],
                                ),
                              ),
                              Padding(
                                padding: EdgeInsets.only(
                                  top:
                                      MediaQuery.sizeOf(context).height * 0.015,
                                ),
                              ),
                              AuthButton(
                                buttonText: "Login",
                                isLoading: loginController.isLoading,
                                onPressed: loginController.isLoading
                                    ? null
                                    : () async {
                                        await loginController.login();
                                        // Show toast after login attempt
                                        Future.microtask(() {
                                          if (loginController.errorMessage !=
                                              null) {
                                            _showToast(
                                                context,
                                                loginController.errorMessage!,
                                                true);
                                          } else if (loginController
                                                  .successMessage !=
                                              null) {
                                            _showToast(
                                                context,
                                                loginController.successMessage!,
                                                false);
                                            Navigator.pushReplacementNamed(
                                                context, '/create-alert');
                                          }
                                        });
                                      },
                              ),
                              const Padding(
                                padding: EdgeInsets.only(top: 20),
                                child: SignUpAndLoginText(isLogin: true),
                              ),
                            ],
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            ),
          );
        },
      ),
    );
  }
}
