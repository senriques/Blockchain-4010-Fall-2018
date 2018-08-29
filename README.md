# University of Wyoming, Blockchain 4010 and 5010

Class at UW (University of Wyoming) for blockchain.

Class 4010 and 5010.

Time: 12:00 (noon) till 12:50 PM, Monday, Wednesday, Friday.

Office Hours:

Tuesday 12:00 to 1:30 PM 

Thursday 5:00 to 6:30 PM

Lecture 01 - Aug 29, 2018
--------------------------

Overview
-----------------

### Lecture 01 - Aug 29 - Wednesday - Class Overview

0. Class Overview
	1. What is Blockchain / Bitcoin - why it is important. 

		In 2009, a person or group of people named Satoshi Nakamoto published “Bitcoin: A Peer-to-Peer Electronic Cash
		System”. ... The Bitcoin design was revolutionary — it elegantly tied
		cryptography, game theory, and economics into a trustless solution to the double-spend problem, and introduced the world
		to the first “chain of blocks”, a censorship-resistant public ledger protected by proof-of-work.

		This is a big deal. Unlike traditional payments, Bitcoin transactions don’t rely on a trusted third-party. Anyone can
		connect to the network and transact, without fear of censorship. Satoshi’s work solved these problems, and founded the
		field of cryptoeconomics.

		In 2013, Vitalik Buterin proposed a new cryptocurrency — Ethereum. Ethereum was Vitalik’s answer to Bitcoin’s
		poor scripting capabilities. Instead of focusing on financial transactions and their outputs, Ethereum transactions are
		about state: agreeing on a computed state, and transitioning from one state to the next.

		Each transaction in Ethereum includes a sender, recipient, funds, and data, similar enough to Bitcoin. Unlike Bitcoin,
		however, a recipient can be a user or a smart contract.

	2. Gartner group projects that 3% of the world economy will blockchain based in 10 years.  This is a compounded annual
		growth rate of 62.2%.

	3. Plan - do lectures in advance of when assignments are due on the material - give students time to do homework.
	   Mark what is going to be tested on.

	4. This class is not a "heavy" programming class.  Yes you will program - but not a huge amount.   Unlike a lot of
		computer science classes this class has a paper and will have test questions involving definitions.  We are
		going to cover some finance, accounting, economics and other topics - not just "how to build a better program".
		If you have a limited programming background I will work with you.

0. What this class will cover
	1. What is Blockchain - what is Bitcoin / Ethereum / Other token systems
	2. The worlds worst, most expensive database
	3. What is the "hype" - what is real about blockchain.
	4. Economics - Coin, ICO, Stocks, Bonds, Tokens, Utility Tokens, A Security
	5. Legal Ramifications.  ICOs 506(d), Subpart (s)
	6. Programming - 1/2 in go, 1/2 in Solidity (Ethereum) and Web front end (JavaScript/HTML/CSS).
	7. Some Homework
	8. Write a Paper - How will blockchain effect the economy.
	9. 2 tests (Midterm and Final)
	10. Why Go
	11. What is Proof of work
	12. What is Proof of stake
	13. Enough Go to make it through this class (and be able to convincingly tell an employer that you have programmed in Go)
	14. Why Ethereum? Solidity?
	15. dApp - what is that?  What is web3?
	16. A detailed understanding of the security model behind Blockchain
	17. Some advanced stuff on security - distributed computation and public/private keys, distributed key generation.
	18. What is a "tangle" - why is a blockchain called a "chain".
	19. Why is blockchain so slow?
	20. How to explain "blockchain" to people - the 30 second elevator pitch.
	21. How to develop for a blockchain.



1. A little bit of background on the instructor
	1. pschlump@uwyo.edu (or pschlump@gmail.com )
	2. https://github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018
	3. (for emergencies between 7:30AM and 9:00PM) 720-209-7888 (cell) 
	4. This is a "practical" class.
	5. Class Goal - have every student in this class be able to work effectively
		in the Blockchain/Ethereum world.   There are students in this class
		that are from other departments.  The class has been structured to 
		take this into account.
	6. Where the students are and the opportunity they have today.
	7. Realizing Your Dreams.

2. Class policy - UW requires that I talk about cheating.
	stackoverflow.com and Google are your fiends in this class - give credit where
	credit is due.  Warning: stackoverflow.com on ethereum/solidity
	is badly out of date becuse the technology is changing really fast.
	If you copy from the web - I expect a comment and a link (URL) to the
	source of where you got your copy.  I expect an explanation of what and 
	why you are grabbing someting from the web.  I expect an analysis of
	what license the content is under, MIT, GPL2, GPL3, CC, CC-BY etc.
	
3. Late Policy and (3) Late Coupons
	I have a grader - most of the grading will be automated.  You get
 	3 late coupons - I will send them to you in email.  You can use
	them for handing in homework 1 week after the due date.  Late 
	coupons will be tracked on a semi-private Ethereum Blockchian
	with a Smart Contract.	Your tokens, pin, password will be
	given to you in a week or so.

	Application: http://www.2c-why.com/UW-4010-dApp/ (Available Sep. 10)

4. 70% from homework, 40% from tests.  Midterm and Final.
	There will be 2 projects for the class.
	You will have to give a live, in person, demo to the instructor of both of your projects
	and you **must** get them to work.  They are required to pass the class.

5. Class Points Total

	Points available:

	| Points            | Class Item         |
	|------------------:|--------------------|
	| 1400              | Homework total     |
	|  800              | 2 tests            |

	To get a letter grade in the class:

	| Points            | Semester Grade     |
	| ----------------- | ------------------ |
	| 1800 or above     | A                  |
	| 1600 ... 1799     | B                  |
	| 1400 ... 1599     | C                  |
	| 1200 ... 1399     | D                  |
	
	You **must** demonstrate working projects to the instructor to pass the class (no matter how many points you get).
	The 2 projects in the class will be directly from the homework.  Project one is Homework Assignments 02 to 06.
	Project 2 is Homework Assignments 08 to 11.  Homework will be 100 or 200 points each.   Midterm will be 400 points.
	Final will be 400 points and cumulative.  You can expect programming and writen assignments in this class.

	For anybody that just wants to take on a hard project for extra credit see the instructor.  It is hard.
	So think a letter grade for completion of an extra credit project.  Code for extra credit projects will
	be open source under a MIT license.   Also note that there are 2200 points available on a letter grading
	scale of 2000 points.   You have a built-in 200 point extra credit in the homework and tests.

5. textbook: [An Introduction to Programming in Go, pdf for free](https://www.golang-book.com/public/pdf/gobook.0.pdf)
	There are no good books on Ethereum/Soliddity.  Solidity has moved from version 4.12 to 4.27 this year.
	All of the books are out of date.  So.... I will include links in assignments that you are expected to
	read.


7. De-Hype Blockchain
	What is a 
		- blockchain: a decentralized, consensus driven store with identities and economic incentives.
		- stake: (as in Proof-of-Stake) (not steak!) an on-chain identity that holds at least the minimum amount of value required to be a participant in the network.
	1. Successes
		- Show Some current screen caputre images.
		- Forcast of blockchain growth.
	2. Failures
		- Over 1/2 of ICOs are proably fraud
		- Slowest most expensive and inconvinent database in existence.





### Lecture 02 - Aug 31 - Friday

1. Economics of Blockchain and Computer Technology - why this is important
	1. How important is blockchain to the economy. 
	2. A $40 billion dollar example.
		[40% decrease in shipping time](https://venturebeat.com/2018/08/09/ibm-and-maersk-launch-blockchain-to-reduce-shipping-time-and-costs/)
	3. Accounting / Banking
	4. Supply Chain
2. Overview of blockchain, bitcoin, EOS, ripple, lightning and Ethereum
3. Good and bad about Go
4. Go

#### References

[Go](https://docs.google.com/presentation/d/1EwuJhEHR5Trr2aXBPQajZ2Hcoh29tm_LQCpgfrCnuRk/preview?slide=id.g33148270ac_0_143)




### Sep 3 - Labor Day - Day Off


### Future lectures will be added....

Yep - schedule is not set in stone yet.




Assignments and Due Dates
--------------------------------

This is a tentative schedule (Subject to change by the instructor):

| Assignment                     | Due Date           | Points |
|:-------------------------------|:-------------------|:------:|
| AS-01 Learn Go                 | Sep 12             | 100    |
| AS-02 Mining Blocks            | Sep 19             | 100    |
| AS-03 Merkle Trees             | Sep 26             | 200    |
| AS-04 Transactions             | Oct 3              | 100    |
| AS-05 Wallets                  | Oct 10             | 100    |
| AS-06 Smart Contract           | Oct 17             | 100    |
| Paper Due                      | Oct 26             | 100    |
| Midterm and Demo Proj 1.       | Week of Oct 29     | 400    |
| AS-07 Solidity & Truffle       | Nov 7              | 100    |
| AS-08 dApp calls Contract      | Nov 14             | 100    |
| AS-09 Grades Contract          | Nov 21             | 100    |
| AS-10 Web Front End (pt1)      | Nov 28             | 100    |
| AS-11 Web Front End (pt2)      | Dec 7              | 200    |
| Demo Proj. 2 start Dec 5       |                    |        |
| ends Dec 14                    |                    |        |
| Final                          |                    | 400    |


Textbook
------------------

[An Introduction to Programming in Go](https://www.golang-book.com/public/pdf/gobook.0.pdf)  The link is to the free PDF of the text.
The book is out of print.

 

Title IX – Duty to Report
----------------------------
 
The University of Wyoming faculty are committed to helping create a safe learning environment for all students and for
the university as a whole. If you have experienced any form of gender or sex-based discrimination or harassment,
including sexual assault, sexual harassment, relationship violence, or stalking, know that help and support are
available. The University has staff members trained to support survivors in navigating campus life, accessing health and
counseling services, providing academic and housing accommodations, and more. The University strongly encourages all
students to report any such incidents to the University. Please be aware that all University of Wyoming employees,
including student staff, are required to report all Title IX related concerns to the Title IX Coordinator or their
supervisor. This means that if you tell a faculty member about a situation of sexual harassment or sexual violence, or
other related misconduct, the faculty member must share that information with the University’s Title IX Coordinator.
UW’s Title IX Coordinator is Jim Osborn (Manager of Investigations, Equal Opportunity Report and Response). He is
located in Room 320 of the Bureau of Mines Building, and can be reached via email at report-it@uwyo.edu or via phone at
766-5200 or 766-5228. For more information, go to:
[http://www.uwyo.edu/reportit/learn-more/faqs.html](http://www.uwyo.edu/reportit/learn-more/faqs.html) .

