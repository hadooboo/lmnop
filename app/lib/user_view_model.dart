import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:lmnop/model.dart';
import 'package:lmnop/repository.dart';

class UserViewModel with ChangeNotifier {
  late final Repository _repository;

  User? get user => _user;
  User? _user;

  Problem? get problem => _problem;
  Problem? _problem;

  UserViewModel(BuildContext context, String userID) {
    _repository = Repository();
    _init(context, userID);
  }

  Future<void> _init(BuildContext context, String userID) async {
    try {
      _user = await _repository.getUser(userID);
    } catch (e) {
      if (e.toString().contains("not found")) {
        Fluttertoast.showToast(msg: "해당 유저 정보가 없습니다.");
      } else {
        Fluttertoast.showToast(msg: "서버에 문제가 발생했습니다.");
      }
      Navigator.pop(context);
      return;
    }
    notifyListeners();

    try {
      _problem = await _repository.getProblem(userID, _user!.getMedianLevel(), []);
    } catch (e) {
      if (e.toString().contains("invalid parameter")) {
        Fluttertoast.showToast(msg: "요청 파라미터가 잘못되었습니다.");
      } else if (e.toString().contains("not found")) {
        Fluttertoast.showToast(msg: "추천할 문제가 없습니다.");
      } else {
        Fluttertoast.showToast(msg: "서버에 문제가 발생했습니다.");
      }
      return;
    }
    notifyListeners();
  }
}