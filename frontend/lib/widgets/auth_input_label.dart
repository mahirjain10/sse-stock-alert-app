import 'package:flutter/material.dart';

// Widget inputLabel(BuildContext context, String label) {
//   return Text(label,
//       style: TextStyle(fontSize: 15, fontWeight: FontWeight.w900));
// }

class AuthInputLabel extends StatelessWidget {
  final String label;
  const AuthInputLabel({Key? key, required this.label}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Text(label,
        style: TextStyle(fontSize: 15, fontWeight: FontWeight.w900));
  }
}
