import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:frontend/utils/validators/auth_validator.dart';

class InputField extends StatelessWidget {
  final String hintText;
  final String field;
  final bool isVisible;
  final VoidCallback? onToggle;
  final void Function(String field, String? errorMsg) validateForm;
  final Map<String, String?> errorText;
  final TextEditingController? controller;

  const InputField({
    Key? key,
    required this.hintText,
    required this.field,
    required this.isVisible,
    required this.validateForm,
    required this.errorText,
    this.onToggle,
    this.controller,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment:
          CrossAxisAlignment.start, // Aligns error text properly
      children: [
        Container(
          color: Colors.amber,
          height: 40,
          width: MediaQuery.sizeOf(context).width * 0.85,
          child: TextFormField(
            autovalidateMode: AutovalidateMode.onUserInteraction,
            controller: controller,
            obscureText: field == "password" ? !isVisible : false,
            // textAlignVertical: TextAlignVertical.center,
            onChanged: (value) {
              final String? error;
              if (field == "email") {
                error = AuthValidator.validateEmail(value);
              } else if (field == "password") {
                error = AuthValidator.validatePassword(value);
              } else {
                error = AuthValidator.validateName(value);
              }
              print(error);
              validateForm(field, error);
            },
            decoration: InputDecoration(
              isDense: true,
              // helperText: '', // Prevents unnecessary space when no error
              fillColor: Colors.white,
              filled: true,
              hintText: hintText,
              border:
                  OutlineInputBorder(borderRadius: BorderRadius.circular(10)),
              // contentPadding:
              //     const EdgeInsets.symmetric(vertical: 12, horizontal: 12),
              suffixIcon: field == "password"
                  ? IconButton(
                      onPressed: onToggle,
                      icon: SvgPicture.asset(
                        isVisible
                            ? "assets/icons/eye.svg"
                            : "assets/icons/eye_slash.svg",
                        width: 20,
                        height: 20,
                      ),
                    )
                  : null,
            ),
          ),
        ),
        if (errorText[field == "password"
                ? "password"
                : field == "email"
                    ? "email"
                    : "name"] !=
            null)
          Padding(
            padding: const EdgeInsets.only(top: 4), // Align error properly
            child: Text(
              errorText[field == "password"
                  ? "password"
                  : field == "email"
                      ? "email"
                      : "name"]!,
              style: const TextStyle(
                fontSize: 13, // Increase readability
                color: Colors.red,
                fontWeight: FontWeight.w500, // Make it slightly bold
              ),
            ),
          ),
      ],
    );
  }
}
