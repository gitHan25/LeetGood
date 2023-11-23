class Solution {
public:
    bool isValid(string s) {
        stack<char> st;
        int length = s.length();
        
        for(int i = 0; i < length; i++){
            if(s[i] == '(' || s[i] == '{' || s[i] == '['){
                st.push(s[i]);
            } else {
                if(st.empty()) return false;
                char top = st.top();
                st.pop();
                if(s[i] == ')' && top != '(') return false;
                if(s[i] == '}' && top != '{') return false;
                if(s[i] == ']' && top != '[') return false;
            }
        }
        return st.empty();
    }
};
