import 'package:flutter/widgets.dart';

class Dashboard extends StatefulWidget {
  final String dashboardText;

  Dashboard({required this.dashboardText});

  @override
  State<Dashboard> createState() => _DashboardState();
}

class _DashboardState extends State<Dashboard> {
  @override
  Widget build(BuildContext context) {
    final screenWidth = MediaQuery.sizeOf(context).width;
    final screenHeight = MediaQuery.sizeOf(context).height;

    return Container(
      decoration: BoxDecoration(
        border: Border.all(style: BorderStyle.solid),
        borderRadius: BorderRadius.all(Radius.circular(10)),
      ),
      height: screenHeight * 0.10,
      width: screenWidth * 0.45,
      padding: EdgeInsets.only(left: 16.0), // Added left padding
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Flexible(
            child: Text(
              widget.dashboardText,
              style: TextStyle(fontWeight: FontWeight.normal, fontSize: 18),
            ),
          ),
          Flexible(
            child: Text(
              "12",
              style: TextStyle(fontWeight: FontWeight.bold, fontSize: 30),
            ),
          )
        ],
      ),
    );
  }
}
