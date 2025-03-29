import 'package:flutter/material.dart';
import 'package:frontend/screens/dashboard/widgets/dashboard.dart';
import 'package:frontend/screens/dashboard/widgets/dashboard_appbar.dart';
import 'package:frontend/screens/dashboard/widgets/recent_alert_card.dart';

class DashboardPage extends StatefulWidget {
  const DashboardPage();
  @override
  State<DashboardPage> createState() => _DashboardPageState();
}

class _DashboardPageState extends State<DashboardPage> {
  @override
  Widget build(BuildContext context) {
    final screenWidth = MediaQuery.sizeOf(context).width;
    final screenHeight = MediaQuery.sizeOf(context).height;
    return SafeArea(
        child: Scaffold(
      backgroundColor: Color(0xFFF3F4F6),
      appBar: DashboardAppbar(),
      body: SingleChildScrollView(
        child: Column(
          children: [
            Container(
              padding: EdgeInsets.only(top: screenHeight * 0.05),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                children: [
                  Dashboard(dashboardText: "Active Alerts"),
                  Dashboard(dashboardText: "Triggered today"),
                ],
              ),
            ),
            Padding(
                padding: EdgeInsets.fromLTRB(
                    screenWidth * 0.05, screenHeight * 0.03, 0, 0),
                child: const Align(
                  alignment: Alignment.topLeft,
                  child: Text(
                    "Recent Alerts",
                    style: TextStyle(fontWeight: FontWeight.bold, fontSize: 20),
                  ),
                )),
            Column(
              children: [RecentAlertCard()],
            )
          ],
        ),
      ),

      // body: Text("YOO"),
    ));
  }
}
