import 'dart:convert';

abstract class UserModel {
  Map<String, dynamic> toJson();
  String toJsonString();
}


class LoginUserModel implements UserModel{
  final String email;
  final String password;
  final String? name; // Nullable, because it might not be present when sending

  LoginUserModel({required this.email, required this.password, this.name});

  // Convert Object to JSON (serialization)
  @override
  Map<String, dynamic> toJson() {
    return {'email': email, 'password': password};
  }

  // Convert Object to JSON String
  @override
  String toJsonString() {
    return jsonEncode(toJson());
  }

  // Convert JSON response to Object (deserialization)
  factory LoginUserModel.fromJson(Map<String, dynamic> json) {
    return LoginUserModel(
      email: json["email"] as String,
      password: "", // Password is not part of the response
      name: json["name"] as String,
    );
  }

  // Convert JSON String to Object
  factory LoginUserModel.fromJsonString(String jsonString) {
    return LoginUserModel.fromJson(jsonDecode(jsonString));
  }
}



class SignupUserModel implements UserModel{
  final String email;
  final String password;
  final String name; 

  SignupUserModel({required this.email, required this.password, required this.name});

  // Convert Object to JSON Map (serialization)
  @override
  Map<String, dynamic> toJson() {
    return {'name':name,'email': email, 'password': password};
  }

  // Convert Object to JSON String
  @override
  String toJsonString() {
    return jsonEncode(toJson());
  }


}
