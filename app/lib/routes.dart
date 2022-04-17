import 'package:app/landing_page.dart';
import 'package:app/user_page.dart';
import 'package:flutter/material.dart';

final routes = {
  '/': (BuildContext context) => LandingPage(),
  '/landing': (BuildContext context) => LandingPage(),
  '/user': (BuildContext context) => UserPage(),
};