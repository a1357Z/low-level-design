#include<bits/stdc++.h>
using namespace std;

struct User
{
    private:
    string name;
    char symbol;

    public:
    User(){}
    User(string name, char symbol){
        this->name = name;
        this->symbol = symbol;
    }

    string getName(){
        return this->name;
    }

    char getSymbol(){
        return this->symbol;
    }
};

struct Board{

    private:
    int size = 3;
    vector<vector<char>> grid;

    public:
    Board(){}

    Board(int size){
        this->size = size;
        grid = vector<vector<char>> (size+1, vector<char>(size+1, '-'));
    }

    vector<vector<char>> getBoard(){
        return this->grid;
    }

    void setBoard(vector<vector<char>> b){
        this->grid = b;
    }

};

class System{
    private:
    Board  board;
    User  user1,  user2, currentPlayer;

    bool checkGameOver(){
        vector<vector<char>> b = this->board.getBoard();

        int emptyCount = 0;
        for(int i=1;i<b.size();i++){
            for(int j=1;j<b.size();j++){
                if(b[i][j] == '-')emptyCount++;
            }
        }

        if(emptyCount == 0){
            cout<<"Game Over"<<endl;
            return true;
        }

        return false;
    }
    
    bool validateInput(int row, int col){
        vector<vector<char>> b = this->board.getBoard();
        if(b[row][col] == '-')return true;
        return false;
    }

    void makeMove(int row, int col){
        char symbol = this->currentPlayer.getSymbol();
        vector<vector<char>> b = this->board.getBoard();
        b[row][col] = symbol;
        this->board.setBoard(b);
    }

    bool checkWinner(){
        //search for any winner and print the winner and return true;
        vector<vector<char>> b = this->board.getBoard();
        
        //scan all rows
        for(int i=1;i<b.size();i++){
            char t = b[i][1];
            bool found = true;
            for(int j=1;j<b.size();j++){
                if(b[i][j] != t){
                    found = false;break;
                }
            }
            if(found && t!= '-')return true;
        }

        //scan all cols
        for(int i=1;i<b.size();i++){
            char t = b[1][i];
            bool found = true;
            for(int j=1;j<b.size();j++){
                if(b[j][i] != t){
                    found = false;break;
                }
            }
            if(found && t!= '-')return true;
        }

        //scan all diagonals hardcoded for size 3 for now
        if(b[1][1] == b[2][2] && b[2][2] == b[3][3] && b[3][3] != '-')return true;
        if(b[1][3] == b[2][2] && b[2][2] == b[3][1] && b[3][1] != '-')return true;

        return false;
    }

    void togglePlayer(){
        User u = this->currentPlayer;
        User u1 = this->user1, u2 = this->user2;
        if(u.getSymbol() == u1.getSymbol())u = u2;
        else
            u = u1;
        
        this->currentPlayer = u;

    }

    void printBoard(){
        vector<vector<char>> b = this->board.getBoard();
        for(int i=1;i<b.size();i++){
            for(int j=1;j<b.size();j++){
                cout<<b[i][j]<<" ";
            }
            cout<<endl;
        }
    }

    public:
    System(){

    }

    System(User u1, User u2){
        this->user1 = u1;
        this->user2 = u2;
        this->currentPlayer = u1;
        this->board = Board(3);
        printBoard();
    }

    bool processInput(int row, int col){
        //validate the input
        if(!validateInput(row, col)){
            cout<<"Invalid move"<<endl;
            return false;
        }

        //make move
        makeMove(row, col);

        //print the board
        printBoard();

        //check if any winner
        if(checkWinner()){
            cout<<"Player "<<this->currentPlayer.getName()<<" won the game."<<endl;
            return true;
        }

        //toggle the player
        togglePlayer();

        //check game over
        if(checkGameOver())return true;

        return false;
    }

    

};

int myAtoi(string &str)
{
    // Initialize result
    int res = 0;
  
    // Iterate through all characters
    // of input string and update result
    // take ASCII character of corresponding digit and
    // subtract the code from '0' to get numerical
    // value and multiply res by 10 to shuffle
    // digits left to update running total
    for (int i = 0; i<str.length(); ++i)
        res = res * 10 + str[i] - '0';
  
    // return result.
    return res;
}

int main(){

    char s1, s2;
    string p1, p2;
    cin>>s1>>p1;
    cin>>s2>>p2;

    User user1(p1, s1);
    User user2(p2, s2);

    System ticTacToeSystem(user1, user2);

    while(true){
        string row, col;
        cin>>row;

        if(row == "exit")break;
        cin>>col;
        int r = myAtoi(row), c = myAtoi(col);
        bool stop = ticTacToeSystem.processInput(r, c);
        if(stop)break;
    }

}
