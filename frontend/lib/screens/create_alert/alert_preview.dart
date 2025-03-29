import 'package:flutter/material.dart';
import 'package:frontend/screens/create_alert/controller/create_alert_controller.dart';

class AlertPreview extends StatefulWidget {
  final CreateAlertController createAlertController;
  const AlertPreview(this.createAlertController, {super.key});

  @override
  State<AlertPreview> createState() => _AlertPreviewState();
}

const Color previewFontColor = Color(0xFF6B7280);
const Color readyFontColor = Color(0xFF1E40AF);
const Color cardBackground = Colors.white;
const Color readyBackground = Color(0xFFDBEAFE);
const Color boldTextColor = Colors.black;

class _AlertPreviewState extends State<AlertPreview> {
  @override
  Widget build(BuildContext context) {
    return Card(
      color: cardBackground,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
      elevation: 2,
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Row(
          crossAxisAlignment: CrossAxisAlignment.center,
          mainAxisSize: MainAxisSize.min, // Ensure row takes minimum height
          children: [
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                mainAxisSize: MainAxisSize.min, // Takes only required space
                children: [
                  const Text(
                    "Preview",
                    style: TextStyle(
                      fontWeight: FontWeight.bold,
                      color: previewFontColor,
                      fontSize: 20,
                    ),
                  ),
                  const SizedBox(height: 12),
                  Text(
                    widget.createAlertController.alertNameController.text
                            .isNotEmpty
                        ? widget.createAlertController.alertNameController.text
                        : "Enter alert name",
                    style: const TextStyle(
                      fontWeight: FontWeight.bold,
                      fontSize: 16,
                      color: boldTextColor,
                    ),
                  ),
                  const SizedBox(height: 6),
                  RichText(
                    text: TextSpan(
                      style: const TextStyle(
                        fontSize: 14,
                        color: previewFontColor,
                      ),
                      children: [
                        const TextSpan(text: "Alert When "),
                        TextSpan(
                          text: widget.createAlertController.ttmController.text
                                  .isNotEmpty
                              ? widget.createAlertController.ttmController.text
                              : "Enter ticker",
                          style: const TextStyle(
                            fontWeight: FontWeight.bold,
                            color: boldTextColor,
                          ),
                        ),
                        TextSpan(
                          text:
                              " ${widget.createAlertController.condiitionTextMap[widget.createAlertController.fieldValue["condition"]]} ",
                        ),
                        TextSpan(
                          text:
                              "\u{20B9} ${widget.createAlertController.alertPriceController.text.isNotEmpty ? widget.createAlertController.alertPriceController.text : "Enter price"}",
                          style: const TextStyle(
                            fontWeight: FontWeight.bold,
                            color: boldTextColor,
                          ),
                        ),
                      ],
                    ),
                  )
                ],
              ),
            ),
            Container(
              color: Colors.red,
              width: 8,
            ),
            Container(
              padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 6),
              decoration: BoxDecoration(
                color: readyBackground,
                borderRadius: BorderRadius.circular(12),
              ),
              child: const Text(
                "Ready",
                style: TextStyle(
                  color: readyFontColor,
                  fontSize: 14,
                  fontWeight: FontWeight.bold,
                ),
              ),
            )
          ],
        ),
      ),
    );
  }
}
