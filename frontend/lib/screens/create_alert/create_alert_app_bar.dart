import 'package:flutter/material.dart';

class CreateAlertAppBar extends StatelessWidget implements PreferredSizeWidget {
  @override
  Widget build(BuildContext context) {
    return AppBar(
      // backgroundColor: const Color(0xFF1D4ED8), // Matching the button color
      elevation: 0,
      automaticallyImplyLeading: false, // Removes the back button
      title: const Text(
        "Create Stock Alert",
        style: TextStyle(
          fontWeight: FontWeight.w600,
          fontSize: 25,
          color: Colors.black,
        ),
      ),
      centerTitle: true,
    );
  }

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);
}
