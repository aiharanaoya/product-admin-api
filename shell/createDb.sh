#!/bin/sh

mysql -u root << EOF
  create database if not exists product_admin_api;
EOF
