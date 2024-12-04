
class Word: 
    def __init__(self, num, str):
        self.num=num
        self.str=str
def sortByNum(val):
    return val.num


def decode():
    list = []
    file = open('coding_qual_input.txt', 'r')
    while True: 
        line = file.readline()
        if not line: 
            break;
        else:
            strings = line.split(" ", 1)
            list.append(Word(int(strings[0]), strings[1]))

    list.sort(key=sortByNum)
    currentIndex = 0 
    currentAdd = 2
    stringToReturn = ""
    while True: 
        stringToReturn = stringToReturn + list[currentIndex].str
        currentIndex=currentIndex+currentAdd
        currentAdd +=1
        if currentIndex >= len(list):
            break

    print(stringToReturn)    
    return(stringToReturn)

decode()