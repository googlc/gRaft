package doc

/*
State:
All servers:
	persistent state:
		currentTerm:
		votedFor
		log[]
	volatile state:
		commitIndex
		lastApplied

Only leader:
	nextIndex[]		for each server,index of the next log entry to send to that server
					(initialized to 0,increases monotonically)
	matchIndex[]	for each server,index of the highest log entry known to be replicated on server
					(initialized to 0,increases monotonically
*/

/*
RPC

*/
