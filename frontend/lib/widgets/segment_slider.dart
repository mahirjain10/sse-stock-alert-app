import 'package:flutter/material.dart';
import 'package:segmented_button_slide/segmented_button_slide.dart';

class SegmentedButtonSlideWidget extends StatefulWidget {
  const SegmentedButtonSlideWidget({Key? key}) : super(key: key);

  @override
  _SegmentedButtonSlideWidgetState createState() =>
      _SegmentedButtonSlideWidgetState();
}

class _SegmentedButtonSlideWidgetState
    extends State<SegmentedButtonSlideWidget> {
  int selectedOption = 0;

  @override
  Widget build(BuildContext context) {
    return Container(
      width: MediaQuery.sizeOf(context).width * 0.90,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            "Condition",
            style: TextStyle(
              fontFamily: 'Poppins',
              fontSize: 16,
              fontWeight: FontWeight.w500,
            ),
          ),
          Padding(padding: EdgeInsetsDirectional.only(bottom: 10)),
          SegmentedButtonSlide(
            selectedEntry: selectedOption,
            onChange: (selected) => setState(() => selectedOption = selected),
            entries: const [
              SegmentedButtonSlideEntry(
                label: ">",
              ),
              SegmentedButtonSlideEntry(
                label: ">=",
              ),
              SegmentedButtonSlideEntry(
                label: "==",
              ),
              SegmentedButtonSlideEntry(
                label: "<",
              ),
              SegmentedButtonSlideEntry(
                label: "<=",
              ),
            ],
            colors: const SegmentedButtonSlideColors(
              barColor: Color(0xFFEFF1F5), // Updated background color to EFF1F5
              backgroundSelectedColor:
                  Color(0xFFFFFFFF), // Selector color set to white
            ),
            slideShadow: [
              BoxShadow(
                color: Colors.grey.withOpacity(1),
                blurRadius: 5,
                spreadRadius: 1,
              )
            ],
            height: 30,
            padding: const EdgeInsets.all(16),
            borderRadius: BorderRadius.circular(12),
            selectedTextStyle: const TextStyle(
              fontWeight: FontWeight.w700,
              color: Colors.black,
              fontSize: 18,
            ),
            unselectedTextStyle: const TextStyle(
              fontWeight: FontWeight.w400,
              fontSize: 15,
              color: Colors.black,
            ),
            hoverTextStyle: const TextStyle(
              fontSize: 15,
              color: Colors.black,
            ),
          ),
        ],
      ),
    );
  }
}
