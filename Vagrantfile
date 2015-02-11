domain = 'example.com'
$script = <<SCRIPT
curl -O http://cbs.centos.org/kojifiles/packages/docker/1.4.1/2.el7/x86_64/docker-1.4.1-2.el7.x86_64.rpm;  
sudo yum localinstall -y docker-1.4.1-2.el7.x86_64.rpm
sudo service docker start
curl -L https://github.com/docker/fig/releases/download/1.0.1/fig-`uname -s`-`uname -m` > /usr/bin/fig; 
chmod +x /usr/bin/fig
fig --version
docker pull bradrydzewski/go:1.3
SCRIPT
 
nodes = [
  { :hostname => 'shipped-dev', :ip => '192.168.99.02', :box => 'msi-centos-7', :ram => 4096, :vcpus => 4},
]

 
 
VAGRANTFILE_API_VERSION = '2'
 
Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  nodes.each do |node|
    config.vm.define node[:hostname] do |node_config|
      node_config.vm.box = node[:box]
      #node_config.vm.box_url = 'http://puppet-vagrant-boxes.puppetlabs.com/ubuntu-server-12042-x64-vbox4210.box'
      node_config.vm.host_name = node[:hostname] + '.' + domain
      node_config.vm.network "private_network", ip: node[:ip] 
 
      vcpus = node.has_key?(:vcpus) ? node[:vcpus] : 1
      memory = node.has_key?(:memory) ? node[:memory] : 512
      disk = node.has_key?(:disk) ? node[:disk] : 0
      size = 512 * 1024
      node_config.vm.provision "shell", inline: $script
 
      node_config.vm.provider :virtualbox do |vb|
        vb.customize ['modifyvm', :id, '--cpus', vcpus]
        vb.customize ['modifyvm', :id, '--memory', memory]
        if node.has_key?(:disk) == true
           for i in 1..disk
               disk_file = "/tmp/" + node[:hostname] + "_" + i.to_s + "_" + rand(1000).to_s +  ".vdi"
               vb.customize ['createhd', '--filename', disk_file, '--size', size]
               vb.customize ['storageattach', :id, '--storagectl', 'SATA Controller', '--port', i, '--device', 0, '--type', 'hdd', '--medium', disk_file]
           end
        end
      end
    end
  end
end
