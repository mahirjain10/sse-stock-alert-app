import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';

class DashboardAppbar extends StatelessWidget implements PreferredSizeWidget {
  const DashboardAppbar({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.red,
      height: 225,
      alignment: Alignment.center,
      padding: const EdgeInsets.symmetric(horizontal: 16),
      child: Row(
        // Use Row for simplicity
        children: [
          const Text(
            "Dashboard",
            style: TextStyle(fontWeight: FontWeight.bold, fontSize: 25),
          ),
          const Spacer(), // Pushes the next Flex to the rightmost side
          Row(
            children: [
              IconButton(
                onPressed: () {},
                icon: SvgPicture.asset(
                  "assets/icons/bell-regular.svg",
                  color: Colors.black,
                  width: 25,
                  height: 25,
                ),
              ),
              Padding(padding: EdgeInsets.only(right: 10)),
              Image.asset(
                "assets/images/profile.png",
                width: 30,
                height: 30,
              ),
            ],
          ),
        ],
      ),
    );
  }

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);
}
