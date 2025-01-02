import 'package:flutter/material.dart';
// import 'package:flutter_svg/flutter_svg.dart';

class AlertHistory extends StatefulWidget {
  const AlertHistory({super.key});
  @override
  State<AlertHistory> createState() => _AlertHistoryState();
}

Color searchBarBackground = const Color(0xFFEFF1F5);

Widget _searchBar() {
  return Container(
    color: searchBarBackground,
    height: 50,
    width: 200,
    child: TextField(
      decoration: InputDecoration(
          border: OutlineInputBorder(), hintText: 'Search  alerts'),
    ),
  );
}

class _AlertHistoryState extends State<AlertHistory> {
  @override
  Widget build(BuildContext context) {
    return Column(
      children: [_searchBar()],
    );
  }
}
