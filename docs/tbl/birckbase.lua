// Generated by github.com/davyxu/tabtoy
// Version: 2.8.10

module table {
export var TBirck : table.ITBirckDefine[] = [
		{ Id : 1, High : 36, Wide : 72 	}
	]


// Id
export var TBirckById : game.Dictionary<table.ITBirckDefine> = {}
function readTBirckById(){
  for(let rec of TBirck) {
    TBirckById[rec.Id] = rec; 
  }
}
readTBirckById();
}
