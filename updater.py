#!/usr/bin/python3

import os
import subprocess as sp

def main():
    accepted_input = ['a', 'b', 'e']
    user_input = input("what should be installed?\n[a]udio, [b]luetooth, [e]verything: ")
    if user_input not in accepted_input:
        print("Please enter [a], [b] or [e]")
        exit(1)
    else:
        update = ['apt', 'update']
        upgrade = ['apt', 'upgrade', '-y']
        prereqs = ['apt', 'install', '-y', 'wget', 'make', 'gcc', 'linux-headers-generic']
        sp.run(update)
        sp.run(upgrade)
        sp.run(prereqs)
        if user_input == 'a':
            install_audio()
        elif user_input == 'b':
            install_bluetooth()
        else:
            install_audio()
            install_bluetooth()
        reboot = ['reboot']
        # sp.run(reboot)

def install_audio():
    print("Installing audio drivers...")
    audio = 'audio/'
    install_cmd_audio = './install.cirrus.driver.sh'
    get_repo_audio = ['git', 'clone', 'https://github.com/jpa-rocha/snd_hda_macbookpro.git', audio]
    sp.run(get_repo_audio)
    os.chdir(audio)
    install_audio = [install_cmd_audio]
    sp.run(install_audio)
    os.chdir('../')
    remove_audio = ['rm', '-rf', audio]
    sp.run(remove_audio)

def install_bluetooth():
    print("Installing bluetooth drivers...")
    bluetooth = 'bluetooth/'
    install_cmd_bluetooth = './install.bluetooth.sh'
    get_repo_bluetooth = ['git', 'clone', 'https://github.com/jpa-rocha/macbook12-bluetooth-driver.git', bluetooth]
    sp.run(get_repo_bluetooth)
    os.chdir(bluetooth)
    install_blue = [install_cmd_bluetooth]
    sp.run(install_blue)
    os.chdir('../')
    remove_bluetooth = ['rm', '-rf', bluetooth]
    sp.run(remove_bluetooth)
    
if __name__ == '__main__':
    main()