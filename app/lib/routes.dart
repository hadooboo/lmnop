import 'package:lmnop/landing_page.dart';
import 'package:lmnop/user_page.dart';
import 'package:flutter/material.dart';

final routes = {
  '/': (BuildContext context) => const LandingPage(),
  '/landing': (BuildContext context) => const LandingPage(),
  '/user': (BuildContext context) => const UserPage(),
};