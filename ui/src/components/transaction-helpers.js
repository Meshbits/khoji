import {secondsToString} from './time-helpers';

export const TX_TYPES = {
  minerReward: 'Miner Reward',
  valueTransfer: 'Transfer',
};

export const transformTransactions = transactions => {
  for (let i = 0; i < transactions.length; i++) {
    transactions[i].timestampHumanReadable = secondsToString(transactions[i].timestamp);
    transactions[i].shielded = transactions[i].shielded ? 'Yes' : 'No';
    transactions[i].type = TX_TYPES[transactions[i].type] ? TX_TYPES[transactions[i].type] : transactions[i].type;
  
    let fundsMovent = {
      in: {},
      out: {},
    };
    console.log('tx data', transactions[i]);

    for (let j = 0; j < transactions[i].vin.length; j++) {
      const vin = transactions[i].vin[j];
      console.log(vin)

      if ('address' in vin) {
        if (!fundsMovent.in[vin.address]) fundsMovent.in[vin.address] = vin.value;
        fundsMovent.in[vin.address] += vin.value;
      }
    }

    for (let j = 0; j < transactions[i].vout.length; j++) {
      const vout = transactions[i].vout[j];
      console.log(vout)

      if ('scriptPubKey' in vout && 'addresses' in vout.scriptPubKey) {
        if (!fundsMovent.out[vout.scriptPubKey.addresses[0]]) fundsMovent.out[vout.scriptPubKey.addresses[0]] = vout.value;
        fundsMovent.out[vout.scriptPubKey.addresses[0]] += vout.value;
        transactions[i].vout[j].address = vout.scriptPubKey.addresses[0];
      }
    }

    transactions[i].movement = {
      input: Object.entries(fundsMovent.in),
      output: Object.entries(fundsMovent.out)
    };

    console.warn('tx ', i, 'funds movement', transactions[i].movement);
  }

  return transactions;
};