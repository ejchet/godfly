#!/usr/bin/env perl
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.
#
# Generate system call table for DragonFly BSD from master list
# (for example, /usr/src/sys/kern/syscalls.master).

use strict;

my $command = "mksysnum_dragonflybsd.pl " . join(' ', @ARGV);

print <<EOF;
// $command
// MACHINE GENERATED BY THE ABOVE COMMAND; DO NOT EDIT

package syscall

const (
EOF

while(<>){
	if(/^([0-9]+)\s+\S+\s+STD\s+({ \S+\s+(\w+).*)$/){
		my $num = $1;
		my $proto = $2;
		my $name = "SYS_$3";
		$name =~ y/a-z/A-Z/;

		# There are multiple entries for enosys and nosys, so comment them out.
		if($name =~ /^SYS_E?NOSYS$/){
			$name = "// $name";
		}
		if($name eq 'SYS_SYS_EXIT'){
			$name = 'SYS_EXIT';
		}

		print "	$name = $num;  // $proto\n";
	}
}

print <<EOF;
)
EOF
