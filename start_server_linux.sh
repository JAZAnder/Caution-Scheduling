#! /bin/sh
while true
do
    printf "\n\nRunning Containers : \n"
    docker ps --format "table {{.ID}}\t{{.Names}}\t{{.Ports}}"
    printf '\nMenu :\n
            1) Run Normal\n
            2) First Time Run (Will Reset Database if exist)\n
            3) Run with PhpMyAdmin\n
            4) First Time Run with PhpMyAdmin (Will Reset Database if exist)\n
            5) Run in Debug Mode (TBA)\n
            6) Kill all Open Containers (TBA)\n
            7) CURL Libraies (Endpoints with Curl) (TBA)\n
            0) Exit Script\n\n'
    
    read -p "Enter your Option: " option

    if [ "$option" = "0" ]; then
        clear
        break
    elif [ "$option" = "1" ];
        then
        #Run Normal
        cp ./config/compose-enviroments/normal-run/docker-compose.yml ./
        docker-compose up -d
        echo ""
        echo "The Website will be avaible on http://local.techwall.xyz"
        echo "Give the container about 30 seconds to get Ready"
        echo -n "Press [ENTER] to continue...: "
        read ignore
        clear
        echo ""
        echo "The Website is avaible on http://local.techwall.xyz"
        echo ""

    elif [ "$option" = "2" ];
        then
        #First Time Run
        cp ./config/compose-enviroments/first-run/docker-compose.yml ./
        docker-compose up -d
        echo ""
        echo "The Website will be avaible on http://local.techwall.xyz"
        echo "Give the container about 30 seconds to get Ready"
        echo -n "Press [ENTER] to continue...: "
        read ignore
        clear
        echo ""
        echo "The Website is avaible on http://local.techwall.xyz"
        echo ""


    elif [ "$option" = "3" ]; 
        then
        #Run with PhpMyAdmin
        cp ./config/compose-enviroments/normal-php/docker-compose.yml ./
        docker-compose up -d
        echo ""
        echo "The Website will be avaible on http://local.techwall.xyz"
        echo "The PhpMyAdmin page will be avaible on http://localhost:8080"
        echo "Give the container about 30 seconds to get Ready"
        echo -n "Press [ENTER] to continue...: "
        read ignore
        clear
        echo ""
        echo "The Website is avaible on http://local.techwall.xyz"
        echo "The PhpMyAdmin page is avaible on http://localhost:8080"
        echo ""


    elif [ "$option" = "4" ]; 
        then
        #First Time Run with PhpMyAdmin
        cp ./config/compose-enviroments/first-php/docker-compose.yml ./
        docker-compose up -d
        echo ""
        echo "The Website will be avaible on http://local.techwall.xyz"
        echo "The PhpMyAdmin page will be avaible on http://localhost:8080"
        echo "Give the container about 30 seconds to get Ready"
        echo -n "Press [ENTER] to continue...: "
        read ignore
        clear
        echo ""
        echo "The Website is avaible on http://local.techwall.xyz"
        echo "The PhpMyAdmin page is avaible on http://localhost:8080"
        echo ""

    
    elif [ "$option" = "5" ]; 
        then
        #Run in Debug Mode
        echo "Feature Coming Soon"
        echo ""
        echo -n "Press [ENTER] to continue...: "
        read ignore
        clear


    elif [ "$option" = "6" ]; 
        then
        #Kill all Open Containers
        docker-compose down -v
        rm ./docker-compose.yml
        echo ""
        echo -n "Press [ENTER] to continue...: "
        read ignore
        clear


    elif [ "$option" = "7" ]; 
        then
        #CURL Libraies
        echo "Feature Coming Soon"
        echo ""
        echo -n "Press [ENTER] to continue...: "
        read ignore
        clear


    fi
done