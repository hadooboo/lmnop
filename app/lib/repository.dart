import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:lmnop/model.dart';

class Repository {
  static const String baseURL = "https://hadooboo-lmnop.herokuapp.com";

  Future<User> getUser(String userID) async {
    final url = Uri.parse(baseURL + "/api/v1/users/$userID");
    final response = await http.get(url);
    switch (response.statusCode) {
      case 200:
        return User.fromJson(json.decode(response.body));
      case 404:
        throw Exception("failed to get user: not found");
      default:
        throw Exception("failed to get user: server error");
    }
  }

  Future<Problem> getProblem(String userID, int level, List<int> except) async {
    final url = Uri.parse(baseURL + "/api/v1/optimum-problem?handle=$userID&level=$level&except=${except.join(",")}");
    final response = await http.get(url);
    switch (response.statusCode) {
      case 200:
        return Problem.fromJson(json.decode(response.body));
      case 401:
        throw Exception("failed to get problem: invalid parameter");
      case 404:
        throw Exception("failed to get problem: not found");
      default:
        throw Exception("failed to get problem: server error");
    }
  }
}