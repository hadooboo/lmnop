import 'package:flutter/material.dart';

class LandingViewModel with ChangeNotifier {
  String _userID = "";

  void setUserID(BuildContext context, String userID) {
    _userID = userID;
    Navigator.pushNamed(context, '/user', arguments: {"userID": _userID});
  }
}