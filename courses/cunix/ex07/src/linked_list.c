#include <stdlib.h>
#include <stdio.h>

typedef struct node {
    void *data;
    struct node *next;
} node_t;

node_t *list_create(void *data)
{
    node_t *new_node = malloc(sizeof(node_t));
    new_node->data = data;
    new_node->next = NULL;

    return new_node;
}

void list_destroy(node_t **head, void (*fp)(void *data))
{
    if (*head != NULL) return;
    list_destroy(&((*head)->next), fp);
    (fp)((*head)->data);
    free(*head);
}   

void list_push(node_t *head, void *data)
{
    node_t *new_node = malloc(sizeof(node_t));
    new_node->data = data;
    new_node->next = NULL;
    
    node_t *temp = head;
    while(temp->next != NULL)
    {
        temp = temp->next;
    }    

    temp->next = new_node;
}

void list_unshift(node_t **head, void *data)
{
    node_t *new_node = malloc(sizeof(node_t));
    new_node->data =data;
    new_node->next = *head;
    *head = new_node;
}   

void *list_pop(node_t **head)
{
    node_t *temp = *head;
    while (temp->next->next != NULL)
    {
        temp = temp->next;
    }
    node_t *del_node = temp->next; 
    void *data = del_node->data;
    free(del_node);
    temp->next = NULL;

    return data;
}

void *list_shift(node_t **head)
{
    node_t *old_head = *head;
    *head = (*head)->next;

    void *data = old_head->data;
    free(old_head);
    return data;
}

void *list_remove(node_t **head, int pos)
{
    node_t *temp = *head;

    for (int i = 0; i < pos -1; i++)
    {
        if (!temp->next)
        {
            return NULL;
        }

        temp = temp->next;
    }
    
    node_t *del_node = temp->next;
    temp->next = temp->next->next;

    void *data = del_node->data;
    free(del_node);
    return data;
}

void list_print(node_t *head)
{
    node_t *temp = head;

    while (temp)
    {
        printf("%s\n", (char *)temp->data);
        temp = temp->next;
    }
}

void list_visitor(node_t *head, void (*fp)(void *data))
{
    node_t *temp = head;

    while (temp)
    {
        (*fp)(temp->data);
        temp = temp->next;
    }
}
