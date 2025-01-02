import 'package:flutter/material.dart';
// import 'package:stock_alert/widgets/SegmentSlider.dart';
import 'package:frontend/widgets/segment_slider.dart';

class Alert {
  final String label;
  final String hintText;

  Alert(this.label, this.hintText);
}

class AlertScreen extends StatefulWidget {
  const AlertScreen({super.key});
  @override
  State<AlertScreen> createState() => _AlertScreenState();
}

Color c1 = const Color(0xFF000000);

Widget _labelAndInputUI(context, String label, String hintText) {
  return Container(
    // color: Colors.deepOrange,
    child: Column(children: [
      SizedBox(
        height: 100,
        width: MediaQuery.of(context).size.width * 0.9,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              label,
              style: const TextStyle(
                  fontFamily: 'Poppins',
                  fontSize: 16,
                  fontWeight: FontWeight.w500),
            ),
            Padding(padding: EdgeInsetsDirectional.only(bottom: 10)),
            SizedBox(
                height: 40,
                child: TextFormField(
                  style: const TextStyle(
                      fontFamily: 'Poppins',
                      fontSize: 14,
                      fontWeight: FontWeight.w300,
                      // backgroundColor: Color(0xFFEDEFF1),
                      color: Color(0xFF2E2E2E)),
                  decoration: InputDecoration(
                    fillColor: Color(0xFFECEFF1),
                    filled: true, // Ensure the fill color is applied
                    border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(
                            12), // Set border radius to 12
                        borderSide: BorderSide.none // No border line
                        ),
                    focusedBorder: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(
                            12), // Set border radius to 12
                        borderSide: BorderSide.none // No border line
                        ), // Remove the border when focused
                    // hintText: "Ex : Tata Motors price below 200",
                    hintText: hintText,
                    contentPadding: const EdgeInsets.symmetric(
                        vertical: 10, horizontal: 10), // Center vertically
                  ),
                ))
          ],
        ),
      )
    ]),
  );
}

class _AlertScreenState extends State<AlertScreen> {
  @override
  Widget build(BuildContext context) {
    List<Alert> alerts = [
      Alert("Alert Name",
          "Ex : Tata Motors price below \u20B9 200"), // Using Unicode
      Alert("Stock Name", "Ex : INFY.NS"),
      Alert("Current Price", "Current Price: \u20B9 100"), // Using Unicode
      Alert("Alert Price", "Alert Price: \u20B9 120"), // Using Unicode
    ];
    return SingleChildScrollView(
      child: Container(
        alignment: Alignment.center,
        color: const Color(0xFFF9FAFB),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            const Padding(padding: EdgeInsetsDirectional.only(top: 10)),
            const Text("Set New Alert",
                style: TextStyle(
                    fontFamily: 'Poppins',
                    fontSize: 26,
                    fontWeight: FontWeight.w600)),
            const Padding(
                padding: EdgeInsetsDirectional.only(top: 20, bottom: 10)),
            Column(
              // Assuming you want to render in a Column
              children: List.generate(alerts.length, (index) {
                // Change 5 to the number of times you want to render
                return _labelAndInputUI(context, alerts[index].label,
                    alerts[index].hintText); // Customize as needed
              }),
            ),
            const SegmentedButtonSlideWidget(),
            const Padding(padding: EdgeInsetsDirectional.only(bottom: 30)),
            SizedBox(
              width: MediaQuery.of(context).size.width * 0.9,
              height: 50, // Set your desired width here
              child: ElevatedButton(
                style: ElevatedButton.styleFrom(
                  backgroundColor: Colors.black,
                  shape: RoundedRectangleBorder(
                    borderRadius:
                        BorderRadius.circular(12.0), // Set border radius here
                  ),
                ),
                onPressed: () {
                  // Action when button is pressed
                },
                child: const Text(
                  "Create Alert",
                  style: TextStyle(
                    fontFamily: 'Poppins',
                    fontWeight: FontWeight.w500,
                    color: Colors
                        .white, // Optional: Set text color for better visibility
                  ),
                ),
              ),
            )
          ],
        ),
      ),
    );
  }
}
