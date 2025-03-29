class CreateAlertValidator {
  static String? validateAlertName(String? alertName) {
    if (alertName == null || alertName.isEmpty) {
      return "Alert name cannot be empty";
    }
    if (alertName.length >= 30) {
      return "Alert name cannot be more than 30 charachter";
    }
    if (alertName.length < 3) {
      return "Alert name cannot be less than 3 charachters";
    }
    return null;
  }

  static String? validateTickerToMonitor(String? ttm) {
    if (ttm == null || ttm.isEmpty) {
      return "Ticker to monitor cannot be empty";
    }
    if (ttm.length >= 10) {
      return "Ticker to monitor cannot be more than 10 charachters";
    }
    return null;
  }

  static String? validateTargetPrice(String? targetPrice) {
    RegExp regex = RegExp(
        r'^\d+(\.\d{1,2})?$'); // Allows numbers with up to 2 decimal places
    if (targetPrice == null || targetPrice.isEmpty) {
      return "Target price cannot be empty";
    }
    if (!regex.hasMatch(targetPrice)) {
      return "Target price should be a valid number (e.g., 100 or 99.99)";
    }
    return null; // Valid case
  }
}
