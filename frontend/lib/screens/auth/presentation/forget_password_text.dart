import 'package:flutter/material.dart';

class ForgetPassword extends StatelessWidget {
  final VoidCallback? onTap;

  const ForgetPassword({Key? key, this.onTap}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onTap ??
          () {
            print("Forgot Password Clicked!");
          },
      child: const Text(
        "Forgot Password?",
        style: TextStyle(
          color: Color(0xff2563eb),
          fontSize: 15,
        ),
      ),
    );
  }
}
