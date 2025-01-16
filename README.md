Bowling Score Calculator
If you are joining tomorrow, we will use ~30 min on fundamentals of getting started with Go using TDD. When learning a new language, tests are important to build confidence that you are on track to mastery.

To focus on Go and testing tomorrow and not the “business rules” Here is a short summary for how to calculate bowling score:

First, read the specification, then think about how you would structure and implement a solution. Don’t be tempted to read one of the countless implementations on github or chatgpt.

Bowling Score Calculation Requirements

Game Structure:

• A game consists of 10 frames.

• Each frame allows up to two rolls unless a strike is achieved.

• If you knock down all 10 pins in:

Two rolls, it is called a spare.

One roll, it is called a strike.

2. Scoring Rules:

Normal frame (neither spare nor strike), the score is the total pins knocked down in both rolls.

Example: If you knock down 1 pin on the first roll and 3 pins on the second roll, the score for the frame is 4.

Spare:

• The score for a spare is 10(for the pins in the frame) + the numberare’s bonus.

• If you score a _strike_ in the 10th frame, you get _2 extra rolls_ to calculate the strike’s bonus.

• Note: No additional points is given to the extra rolls themselves in the 10th frame. They are only there to calculate the value for the strike or spare. of pins knocked down in the next roll.

• Example: If you score a spare in Frame 1 (e.g., 5 pins + 5 pins), and then knock down 3 pins on the first roll of Frame 2, the score for Frame 1 is _13_.

Strike:

• The score for a strike is 10 + the number of pins knocked down in the next two rolls.

• Example: If you score a strike in Frame 1 and then knock down 3 pins and 5 pins in Frame 2, the score for Frame 1 is 18.

3. Special Rules for the 10th Frame:
   • If you score a spare in the 10th frame, you get 1 extra roll to calculate the spare’s bonus.

• If you score a strike in the 10th frame, you get 2 extra rolls to calculate the strike’s bonus.

• Note: No additional points is given to the extra rolls themselves in the 10th frame. They are only there to calculate the value for the strike or spare.

A perfect game is 300 points, as each frame will get the total of 30 (10 for the frame + 10 + 10 for the next two rolls.)
spare’s bonus.

• If you score a strike in the 10th frame, you get 2 extr
