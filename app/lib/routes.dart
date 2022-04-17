import 'package:lmnop/landing_page.dart';
import 'package:lmnop/user_page.dart';
import 'package:flutter/material.dart';

final routes = {
  '/': (BuildContext context) => LandingPage(),
  '/landing': (BuildContext context) => LandingPage(),
  '/user': (BuildContext context) => UserPage(),
};