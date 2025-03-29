import 'package:flutter/material.dart';
import 'package:segmented_button_slide/segmented_button_slide.dart';

class SegmentedButtonSlideWidget extends StatefulWidget {
  final void Function(String field, String? value) setFieldValue;
  const SegmentedButtonSlideWidget({Key? key, required this.setFieldValue})
      : super(key: key);
  @override
  _SegmentedButtonSlideWidgetState createState() =>
      _SegmentedButtonSlideWidgetState();
}

class _SegmentedButtonSlideWidgetState
    extends State<SegmentedButtonSlideWidget> {
  int selectedOption = 0;

  final List<String> conditionList = [">", ">=", "==", "<", "<="];
  @override
  void initState() {
    super.initState();

    // Delay state update until after the first frame
    WidgetsBinding.instance.addPostFrameCallback((_) {
      if (mounted) {
        widget.setFieldValue("condition", conditionList[selectedOption]);
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      width: MediaQuery.sizeOf(context).width * 0.95,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            children: [
              SizedBox(width: 3),
              Text(
                "Condition",
                style: TextStyle(fontWeight: FontWeight.bold, fontSize: 16),
              ),
            ],
          ),
          SizedBox(
            height: 5,
          ),
          SegmentedButtonSlide(
            selectedEntry: selectedOption,
            onChange: (selected) => {
              setState(() => selectedOption = selected),
              widget.setFieldValue("condition", conditionList[selected])
            },
            entries: const [
              SegmentedButtonSlideEntry(label: ">"),
              SegmentedButtonSlideEntry(label: ">="),
              SegmentedButtonSlideEntry(label: "=="),
              SegmentedButtonSlideEntry(label: "<"),
              SegmentedButtonSlideEntry(label: "<="),
            ],
            colors: const SegmentedButtonSlideColors(
              barColor: Color(0xFFE5E7EB), // Light gray to match UI
              backgroundSelectedColor: Color(0xFFFFFFFF), // White selected bg
            ),
            slideShadow: [
              BoxShadow(
                color: Colors.black.withOpacity(0.1), // Softer shadow
                blurRadius: 4,
                spreadRadius: 1,
              )
            ],
            height: 32,
            padding: const EdgeInsets.all(14),
            borderRadius:
                BorderRadius.circular(8), // Less rounded for consistency
            selectedTextStyle: const TextStyle(
              fontWeight: FontWeight.w700,
              color: Color(0xFF1F2937), // Darker text for selection
              fontSize: 16,
            ),
            unselectedTextStyle: const TextStyle(
              fontWeight: FontWeight.w400,
              fontSize: 14,
              color: Color(0xFF6B7280), // Lighter gray for unselected
            ),
            hoverTextStyle: const TextStyle(
              fontSize: 15,
              color: Colors.black,
            ),
          ),
          SizedBox(
            height: 20,
          )
        ],
      ),
    );
  }
}
