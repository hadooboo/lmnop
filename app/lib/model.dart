class User {
  final String handle;
  final int solvedCount;
  final int tier;
  final int rating;
  final int classNumber;
  final String classDecoration;
  final int maxStreak;
  final List<Problem> top100;

  User({
    required this.handle,
    required this.solvedCount,
    required this.tier,
    required this.rating,
    required this.classNumber,
    required this.classDecoration,
    required this.maxStreak,
    required this.top100,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      handle: json['handle'],
      solvedCount: json['solved_count'],
      tier: json['tier'],
      rating: json['rating'],
      classNumber: json['class'],
      classDecoration: json['class_decoration'],
      maxStreak: json['max_streak'],
      top100: json['top_100'].map((problem) => Problem.fromJson(problem)).toList().cast<Problem>(),
    );
  }

  int getMedianLevel() {
    if (top100.isEmpty) {
      return 1;
    }
    return top100[top100.length~/2].level;
  }
}

class Problem {
  final int problemId;
  final int level;

  Problem({required this.problemId, required this.level});

  factory Problem.fromJson(Map<String, dynamic> json) {
    return Problem(
      problemId: json['problem_id'],
      level: json['level'],
    );
  }
}