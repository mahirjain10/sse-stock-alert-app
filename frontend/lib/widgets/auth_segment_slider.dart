import 'package:flutter/material.dart';
import 'package:segmented_button_slide/segmented_button_slide.dart';

class AuthSegmentedButtonSlideWidget extends StatefulWidget {
  const AuthSegmentedButtonSlideWidget({Key? key}) : super(key: key);

  @override
  _AuthSegmentedButtonSlideWidgetState createState() =>
      _AuthSegmentedButtonSlideWidgetState();
}

class _AuthSegmentedButtonSlideWidgetState
    extends State<AuthSegmentedButtonSlideWidget> {
  @override
  Widget build(BuildContext context) {
    int selectedOption = 0;
    return SafeArea(
        child: SegmentedButtonSlide(
      selectedEntry: selectedOption,
      onChange: (selected) => setState(() => selectedOption = selected),
      entries: const [
        SegmentedButtonSlideEntry(
          //   icon: Icons.home_rounded,
          label: "Login",
        ),
        SegmentedButtonSlideEntry(
          icon: Icons.list_rounded,
          label: "Sign Up",
        ),
      ],
      colors: SegmentedButtonSlideColors(
        barColor:
            Theme.of(context).colorScheme.primaryContainer.withOpacity(0.5),
        backgroundSelectedColor: Theme.of(context).colorScheme.primaryContainer,
      ),
    ));
  }
}
