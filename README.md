multithreaded instagram scraper searching for short available usernames based on common patterns.

### available patterns

```c
/*
		- A: firstNameNumSeq
				[firstname] + [numericSequence]
				example: luna221
		- B: firstNameRandWord
				[firstname] + [randomWord]
				example: maxtravels, laurastar
		- C: firstNameRandWordNumSeq
				[PatternB]  + [numericSequence]
				example: maxtravels9, laurastar200
		- D: firstNameDotVerb
				[firstname] + [.] + [verb]
				example: jens.codes, jessie.travels
		- E: randomWordTimesN
				n           * [randomWord]
				example: mytravelhobby, coolmathgames
		- F: randomWordTimesNnumSeq
				[PatternE]  + [numericSequence]
				example: mytravelhobby4, coolmathgames24
		- G: randomPermutation
				[randomPermutation]
				example: aural (laura)
		- H: leetspeak
				[leetspeak(PatternA-G)]
				 example: kyli3j33n3rsb1gg3stst4lk3r
*/
```
