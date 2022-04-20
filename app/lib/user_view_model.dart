import 'package:flutter/material.dart';
import 'package:lmnop/model.dart';
import 'package:lmnop/repository.dart';

class UserViewModel with ChangeNotifier {
  late final Repository _repository;

  User? get user => _user;
  User? _user;

  Problem? get problem => _problem;
  Problem? _problem;

  UserViewModel(String userID) {
    _repository = Repository();
    _init(userID);
  }

  Future<void> _init(String userID) async {
    _user = await _repository.getUser(userID);
    notifyListeners();
    _problem = await _repository.getProblem(userID, _user!.getMedianLevel(), []);
    notifyListeners();
  }
}