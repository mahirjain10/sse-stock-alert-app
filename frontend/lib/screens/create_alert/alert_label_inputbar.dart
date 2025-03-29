import 'dart:async';

import 'package:flutter/material.dart';
import 'package:frontend/screens/create_alert/controller/create_alert_controller.dart';
import 'package:frontend/utils/validators/create_alert_validator.dart';

Color inputBarBorderColor = const Color(0xFFD1D5DB);

class AlertLabelInputbar extends StatefulWidget {
  final String label;
  final String hintText;
  final TextEditingController controller;
  final void Function(String field, String? errorMsg) validateForm;
  final String field;
  final String? errorText;
  final void Function(String field, String? value) setFieldValue;
  final Future<void> Function()? getCurrentPrice; // Fix type

  const AlertLabelInputbar({
    super.key,
    required this.label,
    required this.hintText,
    required this.controller,
    required this.field,
    required this.validateForm,
    required this.setFieldValue,
    this.getCurrentPrice, // Fix nullability
    this.errorText,
  });

  @override
  State<AlertLabelInputbar> createState() => _AlertLabelInputbarState();
}

class _AlertLabelInputbarState extends State<AlertLabelInputbar> {
  Timer? _debounce;
  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Padding(
          padding: const EdgeInsets.only(left: 3),
          child: Text(
            widget.label,
            style: const TextStyle(fontWeight: FontWeight.bold, fontSize: 16),
          ),
        ),
        const SizedBox(height: 5),
        TextFormField(
          keyboardType: widget.field == "alertPrice"
              ? TextInputType.number
              : TextInputType.text,
          controller: widget.controller,
          onChanged: (value) {
            String? error;
            if (widget.field == "alertName") {
              widget.setFieldValue(widget.field, widget.controller.text);
              error = CreateAlertValidator.validateAlertName(value);
            } else if (widget.field == "tickerToMonitor") {
              widget.setFieldValue(widget.field, widget.controller.text);
              error = CreateAlertValidator.validateTickerToMonitor(value);
            } else if (widget.field == "alertPrice") {
              widget.setFieldValue(widget.field, widget.controller.text);
              error = CreateAlertValidator.validateTargetPrice(value);
            } else if (widget.field == "currentPrice") {
              widget.setFieldValue(widget.field, widget.controller.text);
            }
            widget.validateForm(widget.field, error);
            if (widget.field == "tickerToMonitor" && widget.errorText == null) {
              if (_debounce?.isActive ?? false) _debounce!.cancel();
              _debounce = Timer(Duration(milliseconds: 500), () async {
                await widget.getCurrentPrice!();
              });
            }
          },
          // onEditingComplete: () async {
          //   debugPrint("On editing complete");
          //   if (widget.field == "tickerToMonitor") {
          //     debugPrint("here I am ");
          //     await widget.getCurrentPrice!();
          //   }
          // },
          decoration: InputDecoration(
            enabled: widget.field != "currentPrice",
            isDense: true,
            contentPadding:
                const EdgeInsets.symmetric(vertical: 8, horizontal: 10),
            fillColor: Colors.white,
            hintText: widget.hintText,
            border: OutlineInputBorder(
              borderRadius: BorderRadius.circular(8),
              borderSide: BorderSide(color: Colors.grey.shade300, width: 1),
            ),
            focusedBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(8),
              borderSide: BorderSide(color: Colors.blue.shade500, width: 2),
            ),
            enabledBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(8),
              borderSide: BorderSide(color: Colors.grey.shade300, width: 1),
            ),
          ),
        ),
        if (widget.errorText != null) ...[
          const SizedBox(height: 5),
          Padding(
            padding: const EdgeInsets.only(left: 5),
            child: Text(
              widget.errorText!,
              style: const TextStyle(
                fontSize: 12,
                fontWeight: FontWeight.w500,
                color: Colors.red,
              ),
            ),
          ),
        ],
        const SizedBox(height: 20),
      ],
    );
  }
}
