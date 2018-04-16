#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct node
{
    char* current;
    struct node* left;
    struct node* right;
} node_t;
 
node_t* createNode(char* base, int lcount, int rcount, char next);

node_t* createNode(char* base, int lcount, int rcount, char next)
{
    node_t* newNode = (node_t*)malloc(sizeof(node_t));
    static int count = 0;
    
    if (base != NULL)
    {
        // Reset count when starting with new root
        if (strcmp(base, "") == 0)
        {
            count = 0;
        }

        // Create string for this node
        newNode->current = (char*)malloc(strlen(base) + 1);
        strcpy(newNode->current, base);
        newNode->current[strlen(base)] = next;

        // Decrement appropriate count
        if (next == '(')
        {
            lcount--;
        }
        else if (next == ')')
        {
            rcount--;
        }

        // If more right than left, add a node for right )
        if (rcount > lcount)
        {
            newNode->right = createNode(newNode->current, lcount, rcount, ')');
        }

        // If left remaining, create a node for that
        if (lcount > 0)
        {
            newNode->left = createNode(newNode->current, lcount, rcount, '(');
        }
        else if (rcount == 0)
        {
            // Both are zero, terminal case, print it
            count++;
            printf("%d: %s\n", count, newNode->current);
        }
    }
    return newNode;
}

int main()
{
    node_t* root;

    for (int i = 1; i < 8; i++)
    {
        root = createNode("", i, i, '(');
    }
    return 0;
}
