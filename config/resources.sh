#!/bin/bash
# jisnardo@host:~/Downloads/terraform-provider-vsphere-2.2.0$ grep -i resourceV vsphere/provider.go | awk -F ':' '{print $1}' | xargs | sed 's/ /\n/g' > resources.txt

echo "import ("
echo

while read line; do

    ShortGroup=$(echo ${line} | sed 's/_//g')

    echo -e '\t"github.com/crossplane-contrib/provider-jet-vsphere/config/'${ShortGroup}'"'

    d_dir=../config/${ShortGroup}
    
    mkdir -p ${d_dir}

cat <<EOF > ${d_dir}/config.go
package ${ShortGroup}

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("${line}", func(r *config.Resource) {

        // we need to override the default group that terrajet generated for
        // this resource, which would be "github"
        r.ShortGroup = "${ShortGroup}"
    })
}
EOF

done < resources.txt

echo
echo "tjconfig.WithDefaultResourceFn(defaultResourceFn))"
echo

while read line; do

    echo -e '\t\t// "'${line}'$",'

done < resources.txt

echo
echo "for _, configure := range []func(provider *tjconfig.Provider){"
echo

while read line; do

    ShortGroup=$(echo ${line} | sed 's/_//g')

    echo -e "\t\t"${ShortGroup}".Configure,"

done < resources.txt
