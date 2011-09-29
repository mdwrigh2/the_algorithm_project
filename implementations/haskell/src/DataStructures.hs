module TAP.DataStructures where

-- Stack
type Stack a = [a]

stack :: [a] -> Stack a
stack = id

push :: Stack a -> a -> Stack a
push stack item = item:stack

pop :: Stack a -> (a, Stack a)
pop (x:xs) = (x, xs)
pop [] = error "No elements in stack"

peek :: Stack a -> a
peek = head

-- Queue
data Queue a = Queue [a] [a]

-- This takes a list of numbers and turns it into a queue. It assumes the tail
-- of the list is the start of the queue (i.e. it was the first in, and
-- therefore will be the first out)
queue :: [a] -> Queue a
queue xs = Queue [] xs

empty :: Queue a -> Bool
empty (Queue [] []) = True
empty (Queue _ _) = False

qPeek :: Queue a -> a
qPeek (Queue [] [])     = error "No elements in queue!"
qPeek (Queue [] ys)     = qPeek $ Queue (reverse ys) []
qPeek (Queue (x:xs) _ ) = x

enqueue :: Queue a -> a -> Queue a
enqueue (Queue xs ys) y = Queue xs (y:ys)

dequeue :: Queue a -> (a, Queue a)
dequeue (Queue [] [])    = error "No elements in queue!"
dequeue (Queue [] ys)    = dequeue $ Queue (reverse ys) []
dequeue (Queue (x:xs) ys ) = (x, Queue xs ys)

-- Binary tree
data BinTree a = BinTree a (BinTree a) (BinTree a)
               | EmptyTree
               deriving (Show)

-- Performs a sorted insert
insert :: (Ord a) => BinTree a -> a -> BinTree a
insert (EmptyTree) item = BinTree item EmptyTree EmptyTree
insert (BinTree val l r) item | item >= val = BinTree val l $ insert r item
                              | item <  val = BinTree val (insert l item) r

-- Takes a value and deletes it from the tree
delete :: (Ord a) => BinTree a -> a -> BinTree a
delete EmptyTree item = error "Item does not exist in tree"

-- Trie

-- Hash table (association list?)
