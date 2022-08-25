import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

enum Tier {
  bronze, silver, gold, platinum, diamond, ruby, master
}

final Map tierToTierEnum = Map.unmodifiable({
  1: Tier.bronze, 2: Tier.bronze, 3: Tier.bronze, 4: Tier.bronze, 5: Tier.bronze,
  6: Tier.silver, 7: Tier.silver, 8: Tier.silver, 9: Tier.silver, 10: Tier.silver,
  11: Tier.gold, 12: Tier.gold, 13: Tier.gold, 14: Tier.gold, 15: Tier.gold,
  16: Tier.platinum, 17: Tier.platinum, 18: Tier.platinum, 19: Tier.platinum, 20: Tier.platinum,
  21: Tier.diamond, 22: Tier.diamond, 23: Tier.diamond, 24: Tier.diamond, 25: Tier.diamond,
  26: Tier.ruby, 27: Tier.ruby, 28: Tier.ruby, 29: Tier.ruby, 30: Tier.ruby,
  31: Tier.master,
});

final Map tierToTierName = Map.unmodifiable({
  1: "Bronze V", 2: "Bronze IV", 3: "Bronze III", 4: "Bronze II", 5: "Bronze I",
  6: "Silver V", 7: "Silver IV", 8: "Silver III", 9: "Silver II", 10: "Silver I",
  11: "Gold V", 12: "Gold IV", 13: "Gold III", 14: "Gold II", 15: "Gold I",
  16: "Platinum V", 17: "Platinum IV", 18: "Platinum III", 19: "Platinum II", 20: "Platinum I",
  21: "Diamond V", 22: "Diamond IV", 23: "Diamond III", 24: "Diamond II", 25: "Diamond I",
  26: "Ruby V", 27: "Ruby IV", 28: "Ruby III", 29: "Ruby II", 30: "Ruby I",
  31: "Master",
});

final Map tierEnumToTextStyle = Map.unmodifiable({
  Tier.bronze: const TextStyle(color: Color.fromRGBO(173, 86, 0, 1.0)),
  Tier.silver: const TextStyle(color: Color.fromRGBO(67, 95, 122, 1.0)),
  Tier.gold: const TextStyle(color: Color.fromRGBO(236, 154, 0, 1.0)),
  Tier.platinum: const TextStyle(color: Color.fromRGBO(39, 226, 164, 1.0)),
  Tier.diamond: const TextStyle(color: Color.fromRGBO(0, 180, 252, 1.0)),
  Tier.ruby: const TextStyle(color: Color.fromRGBO(255, 0, 98, 1.0)),
  Tier.master: createTextStyleWithColors([
    const Color.fromRGBO(255, 124, 168, 1.0),
    const Color.fromRGBO(180, 145, 255, 1.0),
    const Color.fromRGBO(124, 249, 255, 1.0),
  ]),
});

enum Class {
  c1, c2, c3, c4, c5, c6, c7, c8, c9, c10
}

final Map classToClassEnum = Map.unmodifiable({
  1: Class.c1, 2: Class.c2, 3: Class.c3, 4: Class.c4, 5: Class.c5,
  6: Class.c6, 7: Class.c7, 8: Class.c8, 9: Class.c9, 10: Class.c10,
});

final Map classToClassName = Map.unmodifiable({
  1: "Class 1", 2: "Class 2", 3: "Class 3", 4: "Class 4", 5: "Class 5",
  6: "Class 6", 7: "Class 7", 8: "Class 8", 9: "Class 9", 10: "Class 10",
});

final Map classEnumToTextStyle = Map.unmodifiable({
  Class.c1: createTextStyleWithColors([
    const Color.fromRGBO(73, 251, 254, 1.0),
    const Color.fromRGBO(36, 156, 229, 1.0),
  ]),
  Class.c2: createTextStyleWithColors([
    const Color.fromRGBO(65, 253, 254, 1.0),
    const Color.fromRGBO(32, 197, 233, 1.0),
  ]),
  Class.c3: createTextStyleWithColors([
    const Color.fromRGBO(55, 254, 252, 1.0),
    const Color.fromRGBO(27, 223, 139, 1.0),
  ]),
  Class.c4: createTextStyleWithColors([
    const Color.fromRGBO(88, 253, 69, 1.0),
    const Color.fromRGBO(43, 213, 33, 1.0),
  ]),
  Class.c5: createTextStyleWithColors([
    const Color.fromRGBO(253, 254, 43, 1.0),
    const Color.fromRGBO(176, 219, 21, 1.0),
  ]),
  Class.c6: createTextStyleWithColors([
    const Color.fromRGBO(254, 253, 29, 1.0),
    const Color.fromRGBO(235, 202, 15, 1.0),
  ]),
  Class.c7: createTextStyleWithColors([
    const Color.fromRGBO(255, 253, 38, 1.0),
    const Color.fromRGBO(243, 180, 18, 1.0),
  ]),
  Class.c8: createTextStyleWithColors([
    const Color.fromRGBO(255, 252, 0, 1.0),
    const Color.fromRGBO(255, 125, 0, 1.0),
  ]),
  Class.c9: createTextStyleWithColors([
    const Color.fromRGBO(255, 55, 238, 1.0),
    const Color.fromRGBO(243, 27, 116, 1.0),
  ]),
  Class.c10: createTextStyleWithColors([
    const Color.fromRGBO(253, 67, 255, 1.0),
    const Color.fromRGBO(167, 32, 232, 1.0),
  ]),
});

enum ClassDecoration {
  none, silver, gold
}

final Map classDecorationToClassDecorationEnum = Map.unmodifiable({
  "none": ClassDecoration.none,
  "silver": ClassDecoration.silver,
  "gold": ClassDecoration.gold,
});

final Map classDecorationToClassDecorationName = Map.unmodifiable({
  "none": "",
  "silver": "Silver",
  "gold": "Gold",
});

TextStyle createTextStyleWithColors(List<Color> colors) {
  return TextStyle(
    foreground: Paint()..shader = createShaderWithColors(colors),
  );
}

Shader createShaderWithColors(List<Color> colors) {
  return LinearGradient(
    begin: Alignment.centerLeft,
    end: Alignment.centerRight,
    colors: colors,
  ).createShader(const Rect.fromLTWH(0, 0, 192, 48));
}