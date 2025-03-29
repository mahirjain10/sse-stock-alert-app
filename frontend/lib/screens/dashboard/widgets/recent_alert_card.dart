import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:flutter/widgets.dart';

class RecentAlertCard extends StatefulWidget {
  RecentAlertCard();

  @override
  State<RecentAlertCard> createState() => _RecentAlertCard();
}

class _RecentAlertCard extends State<RecentAlertCard> {
  @override
  Widget build(BuildContext context) {
    final screenWidth = MediaQuery.sizeOf(context).width;
    return Container(
        width: screenWidth * 0.95,
        alignment: Alignment.center,
        child: Card(
            color: Colors.white,
            child: Container(
              color: Colors.red,
              child: Padding(
                padding: EdgeInsets.symmetric(vertical: 5, horizontal: 10),
                child: Column(
                  children: [
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Text(
                          "AAPL Above 180",
                          style: TextStyle(
                              fontWeight: FontWeight.bold, fontSize: 15),
                        ),
                        Text(
                          "Active",
                          style: TextStyle(
                            backgroundColor: Colors.green,
                          ),
                        )
                      ],
                    ),
                    Row(
                      children: [
                        Text(
                          "Apple INC",
                        ),
                      ],
                    ),
                    Row(
                      children: [Text("Current: 242.15"), Text("250")],
                    )
                  ],
                ),
              ),
            )));
  }
}
