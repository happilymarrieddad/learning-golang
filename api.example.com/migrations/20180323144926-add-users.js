'use strict';

var dbm;
var type;
var seed;

/**
  * We receive the dbmigrate dependency from dbmigrate initially.
  * This enables us to not have to rely on NODE_PATH.
  */
exports.setup = function(options, seedLink) {
  dbm = options.dbmigrate;
  type = dbm.dataType;
  seed = seedLink;
};

exports.up = function(db) {
  return db.createTable('users',{
    id: { type: 'int', primaryKey: true, autoIncrement: true },
    first: 'string',
    last:'string',
    email:{type :'string',notNull:true,length:200},
    password:{type :'string',notNull:true,length:200}
  });
};

exports.down = function(db) {
  return db.dropTable('users');
};

exports._meta = {
  "version": 1
};
