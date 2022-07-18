UML diagram: https://miro.com/app/board/uXjVOrwASf8=/

The SignedMessage structure was created to store serialized data and digitally sign it. It has a digital signature verification method.

The KeyPair structure was created to store a key pair, it has methods for generating it and signing data.

The PaymentMessages structure is a container for transaction data, and is designed to serialize and decode this data from a slice of bytes.

The Account and Depository structures are similar, but have radically different functionality. Account has a 256-bit identifier that is a public key. Additionally, it contains the entire key pair as well as the balance. This was done in order to sign the data with the KeyPair structure method and avoid transferring the private key. It has methods for creating a new account from an existing key pair, as well as creating a payment transaction.

The Depository structure was created for the distribution of test coins, and in addition to the fields of the Account structure, it also contains a dictionary of accounts holding a deposit. Receiving a deposit from an account has not yet been worked out, this structure is only used to create the initial distribution of coins across accounts using the SendInterests method.

The Operation structure contains the recipient, the sender, the amount of funds transferred, as well as the payment data signed by the sender. There is duplication, but this simplifies the verification of the operation (there is no need to unpack the signed payment data again). The verification method depends on the Blockchain structure, since for verification, a database with existing accounts is used. The ToByte() method is used to serialize operations into a transaction to calculate the hash of the transaction.

The Transaction structure contains its hash and slice of operations. The method of creating transactions is also dependent on the blockchain due to the verification of transactions submitted to form a transaction. The ToByte() method is used to serialize transactions in a block to calculate the hash of the block.

The Block structure contains the hash of the previous block, its own hash, and a slice of transactions. In the CreateBlock method, transactions are packed and the block hash is calculated.

Blockchain structure contains:

1. CoinDatabase, which is a dictionary of strings representing the HEX representation of the public keys (identifiers) of the accounts associated with their balances.
2. BlockHistory - block slice
3. field Coins - the total number of coins in the system The AddBlock method turned out to be the result of validateBlock, after which it updates the records of all accounts in CoinDatabase, and adds the block to BlockHistory.

The InitBlockchain method takes as input the total number of coins (preferably a multiple of 10,000, not less than 10,000, it is recommended to give 1,000,000 as input), and affects the resonance with the genesis block. Here's how it works:

1. A depository is created, which stores all the coins of the blockchain
2. Depository balance is added in CoinDatabase
3. A first user account is created (genesis account)
4. An operation is created to transfer 0.01% of all coins from the depository to this account
5. A transaction is created consisting of this operation
6. A block is created from this transaction
7. This block is added, if everything is fine, then the number of coins in the system is set, the blockchain is returned with the first recorded block.
