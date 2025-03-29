class AuthValidator {
  static String? validateEmail(String? email) {
    if (email == null || email.isEmpty) {
      return "Email cannot be empty";
    }
    final emailRegex =
        RegExp(r"^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$");
    if (!emailRegex.hasMatch(email)) {
      return "Enter a valid email address";
    }
    return null;
  }

  static String? validatePassword(String? password) {
    if (password == null || password.isEmpty) {
      return "Password cannot be empty";
    }
    if (password.length < 8) {
      return "Password must be at least 8 characters long";
    }
    if (password.length >= 15) {
      return "Password must be less than 15 characters long";
    }
    return null;
  }

  static String? validateName(String? name) {
    if (name == null || name.isEmpty) {
      return "Name cannot be empty";
    }
    if (name.length < 3) {
      return "Name must be at least 3 characters long";
    }
    if (name.length >= 15) {
      return "Password must be less than 15 characters long";
    }
    return null;
  }
}
