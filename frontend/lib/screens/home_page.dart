import 'package:flutter/material.dart';
import 'package:frontend/screens/alert_histroy.dart';
import 'package:frontend/screens/alert_screen.dart';

const tabBarBg = Color(0xFF4A90E2);

class HomePage extends StatefulWidget {
  const HomePage({super.key});
  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage>
    with SingleTickerProviderStateMixin {
  late TabController tabController;
  @override
  void initState() {
    tabController = TabController(length: 2, vsync: this);
    tabController.addListener(() {
      setState(() {});
    });
    super.initState();
  }

  static const List<Widget> _pages = <Widget>[
    AlertScreen(),
    AlertHistory(),
  ];
  @override
  Widget build(BuildContext context) {
    return SafeArea(
        child: Scaffold(
      // appBar: AppBar(
      //   centerTitle: true,
      //   title: const Text("Stock Alert App"),
      // ),
      body: Center(
        child: _pages.elementAt(tabController.index),
      ),
      bottomNavigationBar: BottomNavigationBar(
          selectedItemColor: Colors.white,
          unselectedItemColor: const Color(0xFFD1D5DB),
          backgroundColor: tabBarBg,
          currentIndex: tabController.index,
          onTap: (int index) {
            setState(() {
              tabController.index = index;
            });
          },
          items: [
            BottomNavigationBarItem(
              icon: Image.asset(
                "assets/images/bell.png",
                height: 20,
                width: 20,
                color: Colors.white,
              ),
              label: 'Alert Screen', // Changed 'label' back to 'title'
            ),
            BottomNavigationBarItem(
              icon: Image.asset(
                "assets/images/history.png",
                height: 25,
                width: 25,
                color: Colors.white,
              ),
              label: 'Alert History',
            )
          ]),
    ));
  }
}
