#!/usr/bin/env node
const { execSync } = require('child_process');

const rawMods = execSync('find examples | grep go.mod | sed \'s/go.mod//\'');

const modules = rawMods.toString().split('\n').filter((x) => x !== '');

console.log("::set-output name=template_dir::"+JSON.stringify(modules));
